package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "", "名称")
	flag.IntVar(&age, "age", 0, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "间隔")
	flag.Parse()
	fmt.Println("name=", name, "age=", age, "married", married, "delay=", delay)
	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
}
