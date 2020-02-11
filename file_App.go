package main

import (
	"fmt"
	"os"
)

func file_exists(filepath string) bool {

	f, e := os.Stat(filepath)
	fmt.Println(e)
	if e != nil {
		return false
	}
	if os.IsNotExist(e) {
		return false
	}

	fmt.Print(f.IsDir())
	return true
}
func main() {

	file_exists("./out/d")
}
