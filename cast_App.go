package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	t := cast.ToInt("123")
	fmt.Println(t)
	t = cast.ToInt("ft45")
	fmt.Println(t)
}
