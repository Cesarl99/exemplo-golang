package main

import (
	"fmt"
)

///array
var x [5]int

func main() {
	x[0] = 1
	x[1] = 25
	fmt.Println(x[0], x[1])
	fmt.Println(x)
	fmt.Printf("%t/n", x)
	fmt.Printf("%t", x)
	fmt.Println(len(x))
}
