package auth

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/lopo1/mall/auth/utils"
	"github.com/lopo1/mall/shard/auth/token"
	"github.com/lopo1/mall/shard/id"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"os"
	"strings"
)

const (
	// ImpersonateAccountHeader defines the header for account
	// id impersonation.
	ImpersonateAccountHeader = "impersonate-account-id"
	authorizationHeader      = "authorization"
	bearerPrefix             = "Bearer "
)
var noAuthList = []string{"/auth.User/Login"}
// Interceptor creates a grpc auth interceptor.
func Interceptor(publicKeyFile string,ctx context.Context) (context.Context ,error)  {
	f, err := os.Open(publicKeyFile)
	if err != nil {
		return ctx,fmt.Errorf("cannnot open public key file: %v", err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return ctx,fmt.Errorf("cannot read public key: %v", err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		return  ctx,fmt.Errorf("canot parse public key: %v", err)
	}
	i := &interceptor{
		verifier: &token.JWTTokenVerifier{
			PublicKey: pubKey,
		},
	}
	return i.HandleReq(ctx)
}

type tokenVerifier interface {
	Verify(token string) (*utils.CustomClaims, error)
}

type interceptor struct {
	verifier tokenVerifier
}

func (i *interceptor) HandleReq(ctx context.Context)  (context.Context,error){

	tkn, err := tokenFromContext(ctx)
	if err != nil {
		return  ctx,status.Error(codes.Unauthenticated, "")
	}

	cla, err := i.verifier.Verify(tkn)
	if err != nil {
		return  ctx,status.Error(codes.Unauthenticated, err.Error())
	}
	//account := accountKey{
	//	ID: cla.ID,
	//	NickName: cla.NickName,
	//	Mobile: cla.NickName,
	//	AuthorityId: int(cla.AuthorityId),
	//}
	//aid = cla.NickName
	fmt.Println("cla",cla)
	return ContextWithAccount(ctx,id.AccountID(cla.ID)),nil
}

func impersonationFromContext(c context.Context) string {
	m, ok := metadata.FromIncomingContext(c)
	if !ok {
		return ""
	}

	imp := m[ImpersonateAccountHeader]
	if len(imp) == 0 {
		return ""
	}
	return imp[0]
}

func tokenFromContext(c context.Context) (string, error) {
	unauthenticated := status.Error(codes.Unauthenticated, "")
	m, ok := metadata.FromIncomingContext(c)
	if !ok {
		return "", unauthenticated
	}

	tkn := ""
	for _, v := range m[authorizationHeader] {
		if strings.HasPrefix(v, bearerPrefix) {
			tkn = v[len(bearerPrefix):]
		}
	}
	if tkn == "" {
		return "", unauthenticated
	}

	return tkn, nil
}

type accountKey struct{
}

// ContextWithAccount creates a context with given account.
func ContextWithAccount(c context.Context, aid id.AccountID) context.Context {
	return context.WithValue(c, accountKey{}, aid)
}

func GrpcCheckAuth(publicKeyFile string) grpc.UnaryServerInterceptor {
	isCheckAuth:=false
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		for _,val:=range noAuthList{
			if val==info.FullMethod{
				isCheckAuth = true
				continue
			}
		}
		if isCheckAuth {
			return  handler(ctx, req)
		}
		if publicKeyFile=="" {
			return nil, status.Error(codes.Unauthenticated, "publicKey not valid")
		}
		authCtx,err:=Interceptor(publicKeyFile,ctx)
		if err!=nil{
			return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("token not valid: %v", err))
		}
		return  handler(authCtx, req)
	}

}

func AccountIDFromContext(c context.Context) (id.AccountID,error)  {
	v := c.Value(accountKey{})
	aid, ok := v.(id.AccountID)
	if !ok {
		return aid, status.Error(codes.Unauthenticated, "")
	}
	return aid, nil
}