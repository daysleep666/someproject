package main

import (
	"context"
	"fmt"
	"log"
	"time"

	p "git.ifengidc.com/go/kv/pb"
	"github.com/daysleep666/someproject/grpc/pb"
	"google.golang.org/grpc"
)

var addr = ":8080"

// var addr = "test0.kv.ifeng.com:80"

func main() {
	for i := 0; i < 1; i++ {
		go LoopTestGetStream(fmt.Sprintf("%v", i))
		go LoopTestPutStream(fmt.Sprintf("%v", i))
		// go LoopTestCommon(fmt.Sprintf("%v", i))
		// go GetStream()
		// go PutSteam()
		// go AllSteam()
	}
	// go LoopTestTryStream()
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
			panic("GetStream:" + err.Error())
			log.Println(err)
			break
		}
		log.Println("GetStream:", aa)
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
			panic("PutSteam:" + err.Error())
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
			data, err := allStr.Recv()
			if err != nil {
				panic("AllSteam Recv:" + err.Error())
			}
			log.Println(data)
		}
	}()

	go func() {
		for {
			err := allStr.Send(&pb.StreamReqData{Data: "ssss"})
			if err != nil {
				panic("AllSteam Send:" + err.Error())
			}
			time.Sleep(time.Second)
		}
	}()

	select {}
}

func LoopTestCommon(id string) {
	conn, err := grpc.Dial(addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(time.Second))
	if err != nil {
		panic(err)
		return
	}
	key := "abc"
	c := p.NewKVClient(conn)
	c.Put(context.TODO(), &p.PutRequest{Key: key, Value: []byte{}})
	for {
		c = p.NewKVClient(conn)
		_, err := c.Get(context.TODO(), &p.GetRequest{Key: key})
		if err != nil {
			panic(err)
			continue
		}
		log.Println(id + ":common success")
	}
}

func LoopTestPutStream(id string) {
	//通过grpc 库 建立一个连接
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	//通过刚刚的连接 生成一个client对象。
	c := pb.NewGreeterClient(conn)
	for {
		//调用服务端推送流
		ct, cancel := context.WithCancel(context.Background())
		putRes, err := c.PutStream(ct)
		if err != nil {
			fmt.Println("PutSteam--------------------------------------------------------")
			return
		}
		t := time.After(time.Second * 5)
		for {
			// time.Sleep(time.Second)
			select {
			case <-t:
				cancel()
				goto l

			default:
				err := putRes.Send(&pb.StreamReqData{Data: fmt.Sprintf("%v", id)})
				if err != nil {
					panic(id + ":PutStream _ Error Stop")
					cancel()
					goto l
				}
				log.Println(id + ":put success")
			}
		}
	l:
		log.Println(id + ":new stream-------------------------------------")
	}
}

func LoopTestGetStream(id string) {
	//通过grpc 库 建立一个连接
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	//通过刚刚的连接 生成一个client对象。
	c := pb.NewGreeterClient(conn)
	for {
		//调用服务端推送流
		reqstreamData := &pb.StreamReqData{Data: "aaa"}
		ct, cancel := context.WithCancel(context.Background())
		res, err := c.GetStream(ct, reqstreamData)
		if err != nil {
			fmt.Println("GetSteam--------------------------------------------------------")
			return
		}
		t := time.After(time.Second * 5)
		for {
			select {
			case <-t:
				cancel()
				goto l

			default:
				d, err := res.Recv()
				if err != nil {
					panic(id + ":GetStream _ Error Stop" + err.Error())
					cancel()
					goto l
				}
				// log.Println("GetStream:", aa)
				log.Println(id+":get success", d.GetData())
			}
		}
	l:
		// log.Println("new stream-------------------------------------")
	}
}

func LoopTestTryStream() {
	//通过grpc 库 建立一个连接
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	//通过刚刚的连接 生成一个client对象。
	c := pb.NewGreeterClient(conn)
	for {
		//调用服务端推送流
		ct, cancel := context.WithCancel(context.Background())
		_, err := c.PutStream(ct)
		if err != nil {
			fmt.Println("try stream failed")
		} else {
			fmt.Println("try stream success")
		}
		cancel()
		time.Sleep(time.Second)
	}
}
