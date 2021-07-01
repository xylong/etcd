package main

import (
	"context"
	"fmt"
	"go-etcd/util"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func main() {
	host := util.Get("server", "host")
	port := util.Get("server", "port")
	config := clientv3.Config{
		Endpoints:   []string{host + ":" + port},
		DialTimeout: time.Second * 10,
	}
	client, err := clientv3.New(config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	kv := clientv3.NewKV(client)
	ctx := context.Background()
	res, err := kv.Put(ctx, "/services/user", "user2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
