package main

import (
	"flag"
	"fmt"
	"github.com/lopo1/mall/goods/global"
	"github.com/lopo1/mall/goods/handler"
	"github.com/lopo1/mall/goods/initialize"
	"github.com/lopo1/mall/goods/proto"
	"github.com/lopo1/mall/goods/utils"
	"github.com/lopo1/mall/shard/etcd"
	"github.com/lopo1/mall/shard/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"syscall"
)
var (
	addr = flag.String("addr", ":8081", "address to listen")
	IP = flag.String("ip", "0.0.0.0", "ip地址")
	Port = flag.Int("port", 64679, "端口号")
	EtcdAddr    = flag.String("EtcdAddr", "192.168.5.131:12379", "register etcd address") //etcd的地址
	privateKeyFile = flag.String("private_key_file", "handler/private_key.pem", "private key file")
	authPublicKeyFile = flag.String("auth_public_key_file", "../shard/auth/public_key.pem", "public key file for auth")

)
func main() {


	//初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	initialize.InitEs()
	zap.S().Info(global.ServerConfig)

	flag.Parse()
	zap.S().Info("ip: ", *IP)
	if *Port == 0{
		*Port, _ = utils.GetFreePort()
	}
	serverAddr:=fmt.Sprintf("%s:%d", *IP, *Port)
	etcdServer,err := etcd.NewClient(*EtcdAddr)
	if err != nil {
		zap.S().Fatal("etcdServer start error", zap.Error(err))
	}
	go etcdServer.Register(global.ServerConfig.Name, serverAddr, 5)
	zap.S().Info("port: ", *Port)
	go func() {
		server.RunGRPCServer(&server.GRPCConfig{
			Name:   "goods",
			Addr:   serverAddr,
			AuthPublicKeyFile: *authPublicKeyFile,
			//Logger: logger,
			RegisterFunc: func(s *grpc.Server) {
				proto.RegisterGoodsServer(s, &handler.GoodsServer{
					//TokenExpire:    time.Hour*24*30, // 30天过期
					//TokenGenerator: utils.NewJWTTokenGen("mall/goods", privKey),
				})
			},
		})

	}()


	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s :=<-quit
	etcdServer.UnRegister(global.ServerConfig.Name, serverAddr)
	if i, ok := s.(syscall.Signal); ok {
		os.Exit(int(i))
		zap.S().Info("注销成功")
	} else {
		os.Exit(0)
		zap.S().Info("注销失败")
	}
}
