package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	ens := os.Environ()
	for _, v := range ens {
		//fmt.Println( k,"=",v)
		d := strings.Split(v, "=")
		fmt.Println(d[0], "=", d[1])
	}

}
