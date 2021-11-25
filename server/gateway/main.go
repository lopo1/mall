package main

import (
	"context"
	goodspb "github.com/lopo1/mall/goods/proto"
	"go.uber.org/zap"

	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"net/http"
	//"net/textproto"
	"github.com/lopo1/mall/auth/proto"
	"google.golang.org/grpc"
)

var addr = flag.String("addr", ":8080", "address to listen")
var authAddr = flag.String("auth_addr", "localhost:64676", "address for auth service")
var goodsAddr = flag.String("goods_addr", "localhost:64679", "address for goods service")

var EtcdAddr    = flag.String("EtcdAddr", "192.168.5.131:12379", "register etcd address") //etcd的地址

func main() {
	flag.Parse()

	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	//etcdServer,err := etcd.NewClient(*EtcdAddr)
	//if err != nil{
	//	fmt.Println("etcdServer err:",err)
	//}
	//lit,err:=etcdServer.GetService(context.Background(),"auth")
	//if err!=nil{
	//	fmt.Println("lit err:",err)
	//}
	//fmt.Println("lit:",lit[0].Address)
	//mux := runtime.NewServeMux(runtime.WithMarshalerOption(
	//	runtime.MIMEWildcard, &runtime.JSONPb{
	//		EnumsAsInts: true,
	//		OrigName:    true,
	//	},
	//), runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
	//	if key == textproto.CanonicalMIMEHeaderKey(runtime.MetadataHeaderPrefix+auth.ImpersonateAccountHeader) {
	//		return "", false
	//	}
	//	return runtime.DefaultHeaderMatcher(key)
	//}))
	mux := runtime.NewServeMux()
	serverConfig := []struct {
		name         string
		addr         string
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	}{
		{
			name:         "auth",
			addr:         *authAddr,
			registerFunc: proto.RegisterUserHandlerFromEndpoint,
		},
		{
			name:         "goods",
			addr:         *goodsAddr,
			registerFunc: goodspb.RegisterGoodsHandlerFromEndpoint,
		},

	}

	for _, s := range serverConfig {
		err := s.registerFunc(
			c, mux, s.addr,
			[]grpc.DialOption{grpc.WithInsecure()},
		)
		if err != nil {
			zap.S().Fatalf("cannot register service %s: %v", s.name, err)
		}
	}
	http.HandleFunc("/healthz", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("ok"))
	})
	http.Handle("/", mux)
	zap.S().Fatal(http.ListenAndServe(*addr, nil))
}
