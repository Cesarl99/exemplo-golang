package main

import (
	"fmt"
)

var nome = "luiz"

func main() {

	x := 10 + 10
	y := "bom dia"
	z := 0.35

	fmt.Printf("a soma do x:%v,%T\n", x, x)
	fmt.Printf("o valor de y:%v,%T\n", y, y)
	fmt.Printf("o valor z:%v,%T\n", z, z)
	fmt.Printf("nome:%v,%T\n", nome, nome)

}
