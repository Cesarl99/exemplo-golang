package main

import (
	"fmt"
)

func main() {
	x := 4
	y := 3
	if x == 4 || y == 4 {
		fmt.Println("o x é igual a 4 OU y igual a 3")
	}
	if x == 4 && y == 3 {
		fmt.Println("o x é igual a 4 E o y igual a 3")
	}
}
