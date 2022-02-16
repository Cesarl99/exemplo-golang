package main

import (
	"fmt"
)

///Cap. 12 – Funções – 6. Funções anônimas
///Cap. 12 – Funções – 7. Func como expressão
///determinamos a função e no final dela determinamos o valor para dar entrada dentro dela
///muito util para quando vamos utilizar uma so vez
func main() {
	x := 66
	/*
		func(x int) {
			fmt.Println(x, "vezes 11 é:")
			fmt.Println(x * 11)
		}(x)
		x := 66*/

	y := func(x int) {
		fmt.Println(x, "vezes 11 é:")
		fmt.Println(x * 11)
	}
	y(x)
}
