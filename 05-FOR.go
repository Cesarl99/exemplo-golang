package main

import (
	"fmt"
)

func main() {
	///valor de x é 0,"programa" so vai parar enquando o x for menor que zero,
	///e a cada "verificação" acrecentasse +1 no valor de x
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}
}
