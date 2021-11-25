package main

import (
	"flag"
	"fmt"
	"github.com/lopo1/mall/shard/etcd"
	"github.com/lopo1/mall/shard/server"
	"github.com/lopo1/mall/userflv/global"
	"github.com/lopo1/mall/userflv/handler"
	"github.com/lopo1/mall/userflv/initialize"
	"github.com/lopo1/mall/userflv/proto"
	"github.com/lopo1/mall/userflv/utils"

	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"

)
var (
	addr = flag.String("addr", ":8081", "address to listen")
	IP = flag.String("ip", "0.0.0.0", "ip地址")
	Port = flag.Int("port", 64677, "端口号")
	EtcdAddr    = flag.String("EtcdAddr", "192.168.5.131:12379", "register etcd address") //etcd的地址
	privateKeyFile = flag.String("private_key_file", "handler/private_key.pem", "private key file")
	authPublicKeyFile = flag.String("auth_public_key_file", "../shard/auth/public_key.pem", "public key file for auth")

)
func main() {

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

	serverAddr := fmt.Sprintf("%s:%d", *IP, *Port)
	etcdServer,err := etcd.NewClient(*EtcdAddr)
	if err!=nil{
		zap.S().Error("etcd server conn error : ", err)
	}
	go etcdServer.Register(global.ServerConfig.Name, serverAddr, 5)

	go func() {
		server.RunGRPCServer(&server.GRPCConfig{
			Name:   "userflv",
			Addr:   serverAddr,
			AuthPublicKeyFile: *authPublicKeyFile,
			//Logger: logger,
			RegisterFunc: func(s *grpc.Server) {
				proto.RegisterAddressServer(s, &handler.UserOpServer{})
				proto.RegisterMessageServer(s, &handler.UserOpServer{})
				proto.RegisterUserFavServer(s, &handler.UserOpServer{})

			},
		})

	}()


	//服务注册
	//register_client := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	//serviceId := fmt.Sprintf("%s", uuid.NewV4())
	//err = register_client.Register(global.ServerConfig.Host, *Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	//if err != nil {
	//	zap.S().Panic("服务注册失败:", err.Error())
	//}
	zap.S().Debugf("启动服务器, 端口： %d", *Port)

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
