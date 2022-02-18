package main

import (
	"fmt"
)

func main() {
	x := 0
	for {
		if x < 10 {
			fmt.Println("o x é menor que dez")
			x++
		} else {
			fmt.Println("o x é maior que dez!!!!!!!!!")
			break
		}
	}
	fmt.Println("FIM")
}

///neste caso vai rodar o programa ate ser 10 ou maior,
