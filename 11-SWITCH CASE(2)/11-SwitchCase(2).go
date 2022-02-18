package main

import (
	"fmt"
)

var x interface{}

///1-cria a variavel

func main() {
	x = "palmeiras"
	///2-define a variavel e automaticamente o seu tipo
	switch x.(type) {
	///3-define que o tipo da variavel vai ser determinande para ela se encaixar em algumas das cases
	case int:
		fmt.Println("o valor de x e um numero inteiro")
	case float64:
		fmt.Println("o valor de x e um numero float")
	case bool:
		fmt.Println("o valor de x e bool(true/false)")
	default:
		fmt.Println("o valor de x e um conjunto de string")
		///4-cada case com a determinado tipo de variavel e a mensagem que irá aparce e que era se encaixar
	}

}
