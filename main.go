package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/LXJ0000/gee-rpc/client"
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
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)

	client, err := client.Dial("tcp", <-addr)
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer client.Close()

	time.Sleep(time.Second)
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			args := fmt.Sprintf("geerpc req %d", i)
			var reply string
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Println("reply:", reply)
		}()
	}
	wg.Wait()
}
