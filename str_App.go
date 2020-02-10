package main

import (
	"fmt"
	"strings"
)

func main() {

	s := " java "
	fmt.Println(len(strings.TrimSpace(s)))
	fmt.Println(len(strings.Trim(s, " ")))

}
