package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/daysleep666/someproject/grpc/pb"
	"google.golang.org/grpc"
)

// var addr = ":8080"
var addr = "test0.kv.ifeng.com:80"

func main() {
	for i := 0; i < 1; i++ {
		go func() {
			GetStream()
			PutSteam()
			AllSteam()
		}()
	}
	select {}
}

func GetStream() {
	//通过grpc 库 建立一个连接
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	//通过刚刚的连接 生成一个client对象。
	c := pb.NewGreeterClient(conn)
	//调用服务端推送流
	reqstreamData := &pb.StreamReqData{Data: "aaa"}
	res, _ := c.GetStream(context.Background(), reqstreamData)
	for {
		aa, err := res.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(aa)
	}
}

func PutSteam() {
	//通过grpc 库 建立一个连接
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	//通过刚刚的连接 生成一个client对象。
	c := pb.NewGreeterClient(conn)
	//调用服务端推送流
	putRes, _ := c.PutStream(context.Background())
	i := 1
	for {
		i++
		err := putRes.Send(&pb.StreamReqData{Data: fmt.Sprintf("%v", time.Now().Unix())})
		if err != nil {
			log.Println("end", err)
		}
		time.Sleep(time.Second)

	}
}

func AllSteam() {
	//通过grpc 库 建立一个连接
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	//通过刚刚的连接 生成一个client对象。
	c := pb.NewGreeterClient(conn)
	//调用服务端推送流
	allStr, _ := c.AllStream(context.Background())
	go func() {
		for {
			data, _ := allStr.Recv()
			log.Println(data)
		}
	}()

	go func() {
		for {
			allStr.Send(&pb.StreamReqData{Data: "ssss"})
			time.Sleep(time.Second)
		}
	}()

	select {}
}
