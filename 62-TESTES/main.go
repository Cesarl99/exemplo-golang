package main

import "fmt"

///teste como e teste em tabela
func main() {
	x := Soma(36, 65, 542, 98)
	y := mult(10)
	fmt.Println(x)
	fmt.Println(y)
}
func Soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}

	return total
}

func mult(i int) int {
	resultado := i * 36

	return resultado
}
