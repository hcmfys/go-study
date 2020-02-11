package main

import "fmt"

func main() {

	a := 113549
	base := 128
	ds := make([]int, 20)
	i := 0
	for {
		b := a % base
		fmt.Printf("%02X ", b)
		ds[i] = b
		a = a / base
		i++
		if a == 0 {
			break
		}
	}
	ds = ds[:i]

	format := "%02x * 128 ^ %d %s"
	for i := len(ds) - 1; i >= 0; i-- {
		dot := "+"

		if i == 0 {
			dot = ""
			format = "%02x %s"
			fmt.Printf(format, ds[i], dot)
		} else {
			fmt.Printf(format, ds[i], i, dot)
		}

	}
	fmt.Println("=", a)

	for i := len(ds) - 1; i >= 0; i-- {
		if i > 0 {
			fmt.Printf("%02X ", ds[i]|0x80)
		} else {
			fmt.Printf("%02X ", ds[i])
		}

	}

}
