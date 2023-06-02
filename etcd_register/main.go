package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

type ServiceRegister struct {
	etcdClient    *clientv3.Client
	lease         clientv3.Lease                          //租约
	leaseResp     *clientv3.LeaseGrantResponse            //设置租约时间的返回
	canclefunc    func()                                  //租约撤销
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse //租约keepalive相应chan
	key           string                                  //注册的key
}

func NewServiceRegister(endpoints []string, timeNum int64) (*ServiceRegister, error) {
	conf := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 10 * time.Second,
	}

	//client := new(clientv3.Client)
	server := new(ServiceRegister)

	//连接 etcd
	clientTem, err := clientv3.New(conf)
	if nil != err {
		err = fmt.Errorf("etcd conn error:%s", err.Error())
		log.Println(err)
		return nil, err
	}

	server.etcdClient = clientTem

	return server, nil
}

func main() {
	fmt.Println("Hello register")

}
