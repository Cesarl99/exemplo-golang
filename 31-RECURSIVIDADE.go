package main

import (
	"fmt"
)

///Cap. 12 – Funções – 11. Recursividade
///recursividade e utilizado quando queremos que ele fique repetido ate chegar
///e um resultado para "sair" da função
///podemos utilizar loop
func main() {
	fmt.Println(fatorial(5))
}
func fatorial(x int) int {
	if x == 1 {
		return x
	}
	return x * fatorial(x-1)
}
