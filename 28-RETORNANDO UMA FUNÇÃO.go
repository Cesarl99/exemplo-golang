package main

import (
	"fmt"
)

///Cap. 12 – Funções – 8. Retornando uma função
func main() {
	x := retorno()
	y := x(3)
	fmt.Println(y)
	fmt.Println(retorno()(3))
}
func retorno() func(int) int {
	return func(i int) int {
		return i * 30
	}
}
