package main

import "fmt"

func ps(x int) {

	if x > 0 {
		fmt.Print(x, " ")
		ps(x - 1)
	}
}

func ps2(x int) {
	if x > 0 {
		ps2(x - 1)
		fmt.Print(x, " ")
	}
}
func ps3(n int) {
	if n == 0 {
		fmt.Print("我的小鲤鱼")
	} else {
		fmt.Print("抱着")
		ps3(n - 1)
		fmt.Print("的我")
	}
}
func main() {
	ps(5)
	fmt.Println("\n")
	ps2(5)
	fmt.Println("\n")
	ps3(5)
	fmt.Println("\n")
}
