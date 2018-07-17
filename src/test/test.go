package main

import (
	"github.com/coreos/etcd/clientv3"
	"time"
	"fmt"
	"context"
)

func main() {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:[]string{"172.16.91.165:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("connect failed, err:\n", err)
		return
	}
	fmt.Println("connect success!")
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	_, err = cli.Put(context.Background(),"/registry/log/conf2/", "4444")

	// 操作完毕后，取消etcd
	cancel()

	if err!=nil {
		fmt.Println("put failed, err:\t", err)
		return
	}
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)

	resp, err := cli.Get(ctx, "/registry/log/conf2/")

	cancel()
	if err != nil {
		fmt.Println("get failed, err:\t", err)
		return
	}

	for _, ev := range resp.Kvs{
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}

}


