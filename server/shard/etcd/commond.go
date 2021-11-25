package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"strings"
	"time"
)

const schema = "ns"

var Cli *clientv3.Client

// start

// Option is etcd registry option.
type Option func(o *options)

type options struct {
	ctx       context.Context
	namespace string
	ttl       time.Duration
	maxRetry  int
}

// Registry is etcd registry.
type Registry struct {
	opts   *options
	Client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease
}



func NewClient(etcdAddr string) (*Registry,error) {
	registry := new(Registry)
	var err error
	if registry.Client == nil {
		registry.Client,err = clientv3.New(clientv3.Config{
			Endpoints:   strings.Split(etcdAddr, ";"),
			DialTimeout: 15*time.Second,
		})
	}
	return registry,err
}

//func GetAddress(serviceInstances []*ServiceInstance) string {
//	//for _,kv:=range serviceInstances{
//	//
//	//}
//	if serviceInstances[0].Address!=""{
//		addrInfo:=strings.Split(serviceInstances[0].Address,":")
//		if addrInfo[0] == "0.0.0.0"{
//			return
//		}
//	}
//	return ""
//}