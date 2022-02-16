package main

import (
	"fmt"
)

func main() {
	ss := [][]int{
		[]int{1, 3, 8, 9, 5},
		[]int{4, 6, 2, 1, 0},
		[]int{4, 6, 5, 4, 8},
		///podemos ter varias slice dentro de uma só slice
	}
	fmt.Println(ss[2][2])
	///onde o primeiro é qual slice se encontra e o segunda e qual posição ela se encontra
}
