package util

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

type Service struct {
	client *clientv3.Client
}

func NewService() *Service {
	host := Get("server", "host")
	port := Get("server", "port")
	config := clientv3.Config{
		Endpoints:   []string{host + ":" + port},
		DialTimeout: time.Second * 10,
	}
	client, err := clientv3.New(config)
	if err != nil {
		log.Fatalln(err)
	}
	return &Service{client: client}
}

func (s *Service) RegService(id, name, address string) error {
	kv := clientv3.NewKV(s.client)
	keyPrefix := "/services"
	_, err := kv.Put(context.Background(), keyPrefix+id+"/"+name, address)
	return err
}
