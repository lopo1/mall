package initialize

import (
	"fmt"
	"github.com/lopo1/mall/order/global"
	"github.com/lopo1/mall/order/proto"
	"github.com/lopo1/mall/shard/etcd"
	"google.golang.org/grpc/resolver"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"
)


func InitSrvConn(){
	etcdInfo := global.ServerConfig.EtcdConfig
	r := etcd.NewResolver("192.168.5.131:12379")
	resolver.Register(r)
	fmt.Println("r name = " ,r.Scheme())
	goodsConn, err := grpc.Dial(
		//fmt.Sprintf("consul://%s:%d/%s?wait=14s", etcdInfo.Host, etcdInfo.Port, global.ServerConfig.GoodsSrvInfo.Name),
		//fmt.Sprintf("etcd:///%s",  global.ServerConfig.GoodsSrvInfo.Name),
		//fmt.Sprintf(r.Scheme()+"/%s", global.ServerConfig.GoodsSrvInfo.Name),
		fmt.Sprintf(r.Scheme()+"://goods/%s", global.ServerConfig.GoodsSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		//grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【商品服务失败】")
	}
	//var goodsIds []int32
	global.GoodsSrvClient = proto.NewGoodsClient(goodsConn)
	//goodsIds = append(goodsIds,1)
	//goods, err := global.GoodsSrvClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{Id: goodsIds})
	//if err!=nil{
	//	fmt.Println("err = ",err)
	//}
	//fmt.Println("goods",goods)
	//初始化库存服务连接
	invConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", etcdInfo.Host, etcdInfo.Port, global.ServerConfig.InventorySrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		//grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【库存服务失败】")
	}

	global.InventorySrvClient = proto.NewInventoryClient(invConn)
}