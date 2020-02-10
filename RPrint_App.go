package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 1; i <= 1000; i++ {
		fmt.Print("\r doing ", i)
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("\n done")
}
