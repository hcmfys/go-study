package main

import (
	"bufio"
	"crypto/tls"
	"log"
	"net"
	"os"
)

func main() {

	log.SetFlags(log.Lshortfile)
	pwd, _ := os.Getwd()
	p := pwd + "/tls/cert/"
	cer, err := tls.LoadX509KeyPair(p+"cert.pem", p+"key.pem")

	if err != nil {

		log.Println(err)

		return

	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}

	ln, err := tls.Listen("tcp", ":8000", config)

	if err != nil {

		log.Println(err)

		return

	}

	defer ln.Close()

	for {

		conn, err := ln.Accept()

		if err != nil {

			log.Println(err)

			continue

		}

		go handleConnection(conn)

	}

}

func handleConnection(conn net.Conn) {

	defer conn.Close()

	r := bufio.NewReader(conn)

	for {

		msg, err := r.ReadString('\n')

		if err != nil {

			log.Println(err)

			return

		}

		println(msg)

		n, err := conn.Write([]byte("world\n"))

		if err != nil {

			log.Println(n, err)

			return

		}

	}

}
