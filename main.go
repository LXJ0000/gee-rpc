package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/LXJ0000/gee-rpc/codec"
	"github.com/LXJ0000/gee-rpc/server"
)

func startServer(addr chan string) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", lis.Addr())
	addr <- lis.Addr().String()
	server.Accept(lis)
}

func main() {
	addr := make(chan string)
	go startServer(addr)

	conn, _ := net.Dial("tcp", <-addr)
	defer conn.Close()

	time.Sleep(time.Second)
	_ = json.NewEncoder(conn).Encode(server.DefaultOption) // 发送option进行协议交换
	cc := codec.NewGobCodec(conn)
	for i := 0; i < 3; i++ {
		h := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}
		_ = cc.Write(h, fmt.Sprintf("geerpc req %d", h.Seq))
		_ = cc.ReadHeader(h)
		var reply string
		_ = cc.ReadBody(&reply)
		log.Println("reply:", reply)
	}
}
