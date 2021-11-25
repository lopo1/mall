package middlewares

import (
	"context"
	"github.com/lopo1/mall/auth/model"
	"github.com/lopo1/mall/shard/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
var noAuthList = []string{
	"/auth.User/GetUserList",
}
func GrpcCheckAdmin() grpc.UnaryServerInterceptor {
	isAdmin:=false
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		aid,err:=auth.AccountIDFromContext(ctx)
		if err!=nil{
			return nil, status.Error(codes.Unauthenticated, "permission denied")
		}
		for _,val:=range noAuthList{
			if val==info.FullMethod{
				isAdmin = true
				continue
			}
		}
		if !isAdmin {
			return  handler(ctx, req)
		}
		var user model.User
		userInfo,err := user.GetUserInfoById(aid.Uint())
		if err!=nil{
			return nil, err
		}
		if userInfo.Role<2{
			return nil, status.Error(codes.Unauthenticated, "permission denied")
		}

		return nil, status.Error(codes.Unauthenticated, "permission denied")

	}

}