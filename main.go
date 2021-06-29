package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func main() {
	config:=clientv3.Config{
		Endpoints:            []string{"212.129.241.217:2379"},
		DialTimeout: time.Second*10,
	}
	client,err :=clientv3.New(config)
	if err!=nil {
		log.Fatal(err)
	}
	defer client.Close()

	kv:=clientv3.NewKV(client)
	ctx:=context.Background()
	res,err:=kv.Put(ctx,"/services/user","user2")
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
