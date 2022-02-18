package main

import (
	"fmt"
)

var x bool

func main() {
	x = (10 < 1000)
	y := (300 < 250)
	z := (36 < 44)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}
