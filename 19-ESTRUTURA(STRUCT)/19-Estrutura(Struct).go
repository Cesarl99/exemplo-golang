package main

import (
	"fmt"
)

type mercado struct {
	produto  string
	preço    float64
	promoção bool
}

func main() {
	mercado1 := mercado{
		produto:  "arroz",
		preço:    6.50,
		promoção: true,
	}
	mercado2 := mercado{"feijão", 3.00, false}
	mercado3 := mercado{"pão", 2.00, true}
	fmt.Println(mercado1)
	fmt.Println(mercado2)
	fmt.Println(mercado3)
}

/*
 Struct é um tipo de dados composto que nos permite armazenar valores de tipos *diferentes.*
- Seu nome vem de "structure," ou estrutura.
- Declaração: type x struct { y: z }
- Acesso: x.y
- Exemplo: produto,preço,promoção
*/
