package main

import (
	"encoding/json"
	"fmt"
)

///utilizamos a seguinte maneira para realizar a conversão de golang para JSON
///Cap. 16 – Aplicações – 2. JSON marshal (ordenação)
type pessoa struct {
	Nome      string
	Peso      float64
	Idade     int
	Profissão string
	Salario   int
}

func main() {
	cesar := pessoa{"cesar", 62.3, 22, "almoxarife", 2000}
	joaquim := pessoa{"joaquim", 70.6, 32, "motorista", 3000}
	ana := pessoa{"ana", 120.5, 44, "professora", 2650}
	fmt.Println(cesar)
	fmt.Println(joaquim)
	fmt.Println(ana)
	///TRATATIVAS DE ERROS
	c, err := json.Marshal(cesar)
	if err != nil {
		fmt.Println(err)
	}
	j, err := json.Marshal(joaquim)
	if err != nil {
		fmt.Println(err)
	}
	a, err := json.Marshal(ana)
	if err != nil {
		fmt.Println(err)
	}
	///para demostrar a saida escrita utilizamos o string
	fmt.Println(string(c))
	fmt.Println(string(j))
	fmt.Println(string(a))
}
