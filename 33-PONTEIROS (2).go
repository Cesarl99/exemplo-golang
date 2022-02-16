package main

import (
	"fmt"
)

func main() {
	x := 20
	recvalor(x)
	recponteiro(&x)
}

func recvalor(x int) {
	fmt.Println("o valor de x é :", x)
}
func recponteiro(x *int) {
	fmt.Println("o valor de x esta localizado na memoria", x)
}
