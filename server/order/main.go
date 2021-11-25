package main

import (
	"flag"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/lopo1/mall/order/global"
	"github.com/lopo1/mall/order/handler"
	"github.com/lopo1/mall/order/initialize"
	"github.com/lopo1/mall/order/proto"
	"github.com/lopo1/mall/shard/etcd"
	"github.com/lopo1/mall/shard/server"
	"github.com/lopo1/mall/shard/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"os"
	"os/signal"
	"syscall"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)
var (
	IP = flag.String("ip", "0.0.0.0", "ip地址")
	Port = flag.Int("port", 64678, "端口号")
	EtcdAddr    = flag.String("EtcdAddr", "192.168.5.131:12379", "register etcd address") //etcd的地址
	privateKeyFile = flag.String("private_key_file", "handler/private_key.pem", "private key file")
	authPublicKeyFile = flag.String("auth_public_key_file", "../shard/auth/public_key.pem", "public key file for auth")

)
func main() {

	//初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	initialize.InitSrvConn()
	zap.S().Info(global.ServerConfig)

	flag.Parse()
	zap.S().Info("ip: ", *IP)
	if *Port == 0{
		*Port, _ = utils.GetFreePort()
	}

	zap.S().Info("port: ", *Port)

	//初始化jaeger
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
			LocalAgentHostPort:"192.168.5.131:6831",
		},
		ServiceName:"mxshop",
	}

	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	opentracing.SetGlobalTracer(tracer)
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
				proto.RegisterOrderServer(s, &handler.OrderServer{
					//TokenExpire:    time.Hour*24*30, // 30天过期
					//TokenGenerator: utils.NewJWTTokenGen("mall/goods", privKey),
				})
			},
		})

	}()

	//服务注册
	go etcdServer.Register(global.ServerConfig.Name, serverAddr, 5)

	zap.S().Debugf("启动服务器, 端口： %d", *Port)

	//监听订单超时topic
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"192.168.5.131:9876"}),
		consumer.WithGroupName("mall-order"),
	)

	if err := c.Subscribe("order_timeout", consumer.MessageSelector{},handler.OrderTimeout); err != nil {
		fmt.Println("读取消息失败")
	}
	_ = c.Start()
	//不能让主goroutine退出


	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s :=<-quit
	_ = c.Shutdown()
	_ = closer.Close()
	etcdServer.UnRegister(global.ServerConfig.Name, serverAddr)
	if i, ok := s.(syscall.Signal); ok {
		os.Exit(int(i))
		zap.S().Info("注销成功")
	} else {
		os.Exit(0)
		zap.S().Info("注销失败")
	}
}
