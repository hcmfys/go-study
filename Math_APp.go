package main

import (
	"fmt"
	"math"
)

/**
golang math.Floor实现四舍五入
*/
func main() {
	x := 2.0 / 3.0
	a := math.Floor(x)
	b := math.Floor(x + 0.5)
	y := 1.0 / 3
	e := math.Floor(y)
	f := math.Floor(y + 0.5)
	z := 1.0 / 2
	m := math.Floor(z)
	n := math.Floor(z + 0.5)
	fmt.Printf("x=%f,a=%f,b=%f,\ny=%f,e=%f,f=%f,\nz=%f,m=%f,n=%f", x, a, b, y, e, f, z, m, n)
}
