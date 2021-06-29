golang+etcd

### 库
- [etcd](https://github.com/etcd-io/etcd/release)


docker run -it --name etcd -p 2379:2379 -v /home/xyl/etcd:/etcd golang:1.16-alpine sh
docker cp etcd etcd:/usr/bin && docker cp etcdctl etcd:/usr/bin

etcd操作库安装
```bigquery
go get go.etcd.io/etcd/clientv3
// 报错就在mod添加
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
```