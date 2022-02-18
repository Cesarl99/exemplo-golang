package main

import (
	"fmt"
)

///Cap. 14 – Ponteiros – 1. O que são ponteiros?
func main() {
	x := 10
	y := &x
	*y++
	fmt.Println(&x)           ///a localização da variavel x
	fmt.Println(*y)           ///o valor de y tendo como paramentro a localização(ponteiro) de x
	fmt.Printf("%T,%T", x, y) ///o tipo de cada variavel
}
