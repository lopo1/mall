package main

import (
	"flag"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/lopo1/mall/inventory/global"
	"github.com/lopo1/mall/inventory/handler"
	"github.com/lopo1/mall/inventory/initialize"
	"github.com/lopo1/mall/inventory/proto"
	"github.com/lopo1/mall/shard/etcd"
	"github.com/lopo1/mall/shard/server"
	"github.com/lopo1/mall/shard/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"syscall"
)
var (
	addr = flag.String("addr", ":8081", "address to listen")
	IP = flag.String("ip", "0.0.0.0", "ip地址")
	Port = flag.Int("port", 50053, "端口号")
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

	serverAddr:=fmt.Sprintf("%s:%d", *IP, *Port)
	etcdServer,err := etcd.NewClient(*EtcdAddr)
	if err != nil {
		zap.S().Fatal("etcdServer start error", zap.Error(err))
	}
	go etcdServer.Register(global.ServerConfig.Name, serverAddr, 5)
	zap.S().Info("port: ", *Port)
	go func() {
		server.RunGRPCServer(&server.GRPCConfig{
			Name:   "inventory",
			Addr:   serverAddr,
			AuthPublicKeyFile: *authPublicKeyFile,
			//Logger: logger,
			RegisterFunc: func(s *grpc.Server) {
				proto.RegisterInventoryServer(s, &handler.InventoryServer{
					//TokenExpire:    time.Hour*24*30, // 30天过期
					//TokenGenerator: utils.NewJWTTokenGen("mall/goods", privKey),
				})
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
	//zap.S().Debugf("启动服务器, 端口： %d", *Port)

	//监听库存归还topic
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"192.168.5.131:9876"}),
		consumer.WithGroupName("mall-inventory"),
	)

	if err := c.Subscribe("order_reback", consumer.MessageSelector{},handler.AutoReback); err != nil {
		fmt.Println("读取消息失败")
	}
	_ = c.Start()
	//不能让主goroutine退出

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	_ = c.Shutdown()
	//if err = register_client.DeRegister(serviceId); err != nil {
	//	zap.S().Info("注销失败:", err.Error())
	//}else{
	//	zap.S().Info("注销成功:")
	//}
}
