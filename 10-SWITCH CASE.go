package main

import (
	"fmt"
)

var total int

func main() {
	x := 5
	switch {
	case x == 5:
		fmt.Println("x é igual a 5")
	case x < 5:
		fmt.Println("x é manor que 5")
	case x > 5:
		fmt.Println("x é maior que 5")
		fallthrough
	case x == 10:
		fmt.Println("x é igual a 10")
		/*default
		é util quando não se encaixa em nenhum dos caso*/
	}

}
