package main

import (
	"fmt"
	"net"
	"org.springbus/protol"
	"os"
	"sync"
)

func sender(conn net.Conn) {

	defer wg.Done()
	for i := 0; i < 100; i++ {
		words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"
		conn.Write(protocol.Packet([]byte(words)))
	}
	fmt.Println("send over")

}

var wg sync.WaitGroup

func main() {
	server := "127.0.0.1:9988"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("connect success")
	wg.Add(1)
	go sender(conn)

	wg.Wait()
}
