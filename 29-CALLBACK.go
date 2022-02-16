package main

import (
	"fmt"
)
///callback
///o callback ele executa o função necessarios para retornar um valor "completo" para a variavel
func main() {
	t := somenteimpares(soma, []int{51, 52, 53, 54, 56, 55, 57, 58, 59, 60, 61}...)
	fmt.Println(t)
}
func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}
func somenteimpares(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if v%2 != 0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}
