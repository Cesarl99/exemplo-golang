package main

import (
	"fmt"
)

///Desenrolando (enumerando) uma slice
func main() {
	numeros := []int{10, 25, 36, 25, 14, 10}
	total := soma(numeros...) ///para desenrolar uma slice fazemos da seguinte forma
	fmt.Print(total)
}
func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}
