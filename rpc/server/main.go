// 将函数转变为服务
// 手动测试：
// `telnet localhost 1234`
// {"method":"DemoService.Div", "params":[{"A":18,"B":3}], "id":123}

package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	rpcdemo "github.com/zhibailin/go-distributed-crawler-from-scratch/rpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})         // 登记/发布 方法到 Server
	listener, err := net.Listen("tcp", ":1234") // 设置 listener
	if err != nil {
		panic(err)
	}

	// listener 开始监听
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error : %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}

}
