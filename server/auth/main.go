package main

import (
	"flag"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/lopo1/mall/auth/global"
	"github.com/lopo1/mall/auth/handler"
	"github.com/lopo1/mall/auth/initialize"
	"github.com/lopo1/mall/auth/proto"
	auth_utils "github.com/lopo1/mall/auth/utils"
	"github.com/lopo1/mall/shard/utils"
	"github.com/lopo1/mall/shard/etcd"
	"github.com/lopo1/mall/shard/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"
)
var (
	addr = flag.String("addr", ":8081", "address to listen")
	IP = flag.String("ip", "0.0.0.0", "ip地址")
	Port = flag.Int("port", 64676, "端口号")
	EtcdAddr    = flag.String("EtcdAddr", "192.168.5.131:12379", "register etcd address") //etcd的地址
	privateKeyFile = flag.String("private_key_file", "handler/private_key.pem", "private key file")
	authPublicKeyFile = flag.String("auth_public_key_file", "../shard/auth/public_key.pem", "public key file for auth")

)

func main()  {
	//初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	zap.S().Info(global.ServerConfig)
	flag.Parse()
	zap.S().Info("ip: ", *IP)
	if *Port == 0{
		*Port, _ = utils.GetFreePort()
	}
	zap.S().Info("port: ", *Port)
	if global.ServerConfig.Name==""{
		global.ServerConfig.Name = "auth"
	}
	//grpcServer := grpc.NewServer()
	//proto.RegisterUserServer(grpcServer, &handler.UserServer{})
	//lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	//if err != nil {
	//	panic("failed to listen:" + err.Error())
	//}
	//注册服务健康检查
	//grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	//将服务地址注册到etcd中
	serverAddr := fmt.Sprintf("%s:%d", *IP, *Port)
	etcdServer,err := etcd.NewClient(*EtcdAddr)
	go etcdServer.Register(global.ServerConfig.Name, serverAddr, 5)



	pkFile, err := os.Open(*privateKeyFile)
	if err != nil {
		zap.S().Fatal("cannot open private key", zap.Error(err))
	}

	pkBytes, err := ioutil.ReadAll(pkFile)
	if err != nil {
		zap.S().Fatal("cannot read private key", zap.Error(err))
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	if err != nil {
		zap.S().Fatal("cannot parse private key", zap.Error(err))
	}
	go func() {
		server.RunGRPCServer(&server.GRPCConfig{
			Name:   "auth",
			Addr:   serverAddr,
			AuthPublicKeyFile: *authPublicKeyFile,
			//Logger: logger,
			RegisterFunc: func(s *grpc.Server) {
				proto.RegisterUserServer(s, &handler.UserServer{
					TokenExpire:    time.Hour*24*30, // 30天过期
					TokenGenerator: auth_utils.NewJWTTokenGen("mall/auth", privKey),
				})
			},
		})

	}()

	//优雅关闭服务
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