package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/daysleep666/someproject/grpc/pb"
	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) GetStream(req *pb.StreamReqData, res pb.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		err := res.Send(&pb.StreamResData{Data: fmt.Sprintf("%v", time.Now().Unix())})
		if err != nil {
			log.Println("getstream", err)
			break
		}
		// if i > 10 {
		// 	break
		// }
		// time.Sleep(time.Second)
	}
	return nil
}

func (s *server) PutStream(req pb.Greeter_PutStreamServer) error {
	for {
		if tem, err := req.Recv(); err == nil {
			log.Println(tem)
		} else {
			log.Println("break,err:", err)
			break
		}
	}
	return nil
}

func (s *server) AllStream(allStr pb.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			allStr.Recv()
			// log.Println(data)
		}
		wg.Done()
	}()

	go func() {
		for {
			allStr.Send(&pb.StreamResData{Data: "ssss"})
			time.Sleep(time.Second)
		}
		wg.Done()
	}()

	wg.Wait()
	return nil
}

func main() {
	//监听端口
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}
	//创建一个grpc 服务器
	s := grpc.NewServer()

	//注册事件
	pb.RegisterGreeterServer(s, &server{})
	//处理链接
	go s.Serve(lis)

	log.Println("start")
	// select {
	// case <-time.After(time.Second * 5):
	// 	break
	// }
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sg := <-quit
	log.Println("Receive signal %v and shutdown server...", sg)

	fmt.Println("prepare...............")
	st := time.Now().Unix()
	s.GracefulStop()
	a := time.Now().Unix() - st
	// time.Sleep(time.Second * 10)
	log.Println("end", a)
}
