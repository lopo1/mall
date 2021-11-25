package etcd

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)



//将服务地址注册到etcd中
func (r *Registry) Register(serviceName, serverAddr string, ttl int64) error {

	//与etcd建立长连接，并保证连接不断(心跳检测)
	ticker := time.NewTicker(time.Second * time.Duration(ttl))
	go func() {
		key := "/" + schema + "/" + serviceName + "/" + serverAddr
		for {
			resp, err := r.Client.Get(context.Background(), key)
			//fmt.Printf("resp:%+v\n", resp)
			if err != nil {
				fmt.Printf("获取服务地址失败：%s", err)
			} else if resp.Count == 0 { //尚未注册
				err = r.keepAlive(serviceName, serverAddr, ttl)
				if err != nil {
					fmt.Printf("保持连接失败：%s", err)
				}
			}
			<-ticker.C
		}
	}()

	return nil
}
//保持服务器与etcd的长连接
func (r *Registry)keepAlive(serviceName, serverAddr string, ttl int64) error {
	//创建租约
	leaseResp, err := r.Client.Grant(context.Background(), ttl)
	if err != nil {
		fmt.Printf("创建租期失败：%s\n", err)
		return err
	}

	//将服务地址注册到etcd中
	key := "/" + schema + "/" + serviceName + "/" + serverAddr
	_, err = r.Client.Put(context.Background(), key, serverAddr, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		fmt.Printf("注册服务失败：%s", err)
		return err
	}

	//建立长连接
	ch, err := r.Client.KeepAlive(context.Background(), leaseResp.ID)
	if err != nil {
		fmt.Printf("建立长连接失败：%s\n", err)
		return err
	}

	//清空keepAlive返回的channel
	go func() {
		for {
			<-ch
		}
	}()
	return nil
}

//取消注册
func (r *Registry) UnRegister(serviceName, serverAddr string) {
	if r.Client != nil {
		key := "/" + schema + "/" + serviceName + "/" + serverAddr
		r.Client.Delete(context.Background(), key)
	}
}

//rch := cli.Watch(context.Background(), "/logagent/conf/")
//        for wresp := range rch {
//            for _, ev := range wresp.Events {
//                fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
//            }
//        }
type ServiceInstance struct {
	 Address string
}
// GetService return the service instances in memory according to the service name.
func (r *Registry) GetService(ctx context.Context, name string) ([]*ServiceInstance, error) {
	//key := fmt.Sprintf("%s/%s", r.opts.namespace, name)
	key := "/" + schema + "/" + name
	//resp, err := r.kv.Get(ctx, key, clientv3.WithPrefix())
	resp, err := r.Client.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	//fmt.Println("key:",key)
	//resp:= r.Client.Watch(ctx, key)
	//cli.Watch(context.Background(), "/logagent/conf/")

	var items []*ServiceInstance

		for  _,kv := range resp.Kvs {
			//fmt.Printf("%q : %q\n",  kv.Key, kv.Value)
			si := ServiceInstance{Address: string(kv.Value)}
			//fmt.Printf("key:%s\n",kv.Value)
			//fmt.Printf("key:%d\n",kv.Value)
			items = append(items, &si)


		}

	return items, nil
}