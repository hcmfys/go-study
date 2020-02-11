package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

//生成私钥：
//openssl genrsa -out key.pem 2048
//生成证书：
//openssl req -new -x509 -key key.pem -out cert.pem -days 3650

func main() {
	http.HandleFunc("/hello", HelloServer)
	pwd, _ := os.Getwd()
	p := pwd + "/tls/cert/"
	fmt.Println("tls listen port 8080")
	e := http.ListenAndServeTLS(":8080", p+"cert.pem", p+"key.pem", nil)
	if e != nil {
		panic(e)
	}
}
