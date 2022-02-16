package main

import (
	"fmt"
)
///struct

type pessoa struct {
	nome  string
	idade int
}
type profissional struct {
	pessoa
	profissão string
	salario   int
}

func main() {
	pessoa1 := pessoa{
		nome:  "cesar",
		idade: 22,
	}
	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "luiz",
			idade: 30,
		},
		profissão: "motorista",
		salario:   3600,
	}
	pessoa3 := pessoa{"jose", 55}
	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
	fmt.Println(pessoa1.idade)
	fmt.Println(pessoa2.nome)
	fmt.Println(pessoa3.nome)
}
