package main

import (
	"fmt"
)

///podemos "puxar" a função ja definida
func main() {
	primeiro()                                       ///primeira função
	segundo("manhã")                                 ///segunda função
	numeros, qtd := soma(10, 25, 36, 99, 87, 44, 12) ///terceira função
	fmt.Println(numeros, qtd)
}
func primeiro() { ///<-primeira função
	fmt.Println("oi,bom dia !!!!!!!!!!!!!")

}

func segundo(s string) { ///<-segunda função
	if s == "manhã" {
		fmt.Println("tenha um otimo dia !!!!!!!!!!")
	} else if s == "tarde" {
		fmt.Println("tenha uma otima tarde!!!!")
	} else {
		fmt.Println("tenha uma otima noite!!!!!!!!")
	}

}
func soma(x ...int) (int, int) { ///<-terceira função
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x)
}
