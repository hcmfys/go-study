package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
)

func main() {

	t := fsnotify.Chmod.String()
	fmt.Println("api")
	fmt.Println(t)
}
