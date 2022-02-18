package main

import (
	"encoding/json"
	"fmt"
)

///Cap. 16 – Aplicações – 3. JSON unmarshal (desordenação)
type dados struct {
	Nome     string  `json:"Nome"`
	Peso     float64 `json:"Peso"`
	Idade    int     `json:"Idade"`
	ProfissO string  `json:"Profissão"`
	Salario  int     `json:"Salario"`
}

func main() {
	pessoa1 := []byte(`{"Nome":"cesar","Peso":62.3,"Idade":22,"Profissão":"almoxarife","Salario":2000}`)
	///1-utilizar slice dado pelo JSON
	var cesar dados ///2-utilizamos o var ,seguindo da variavel criada para
	///receber os dados,seguindo do nome da estrutura
	err := json.Unmarshal(pessoa1, &cesar) ///3-utilizamos o json.Unmarshal para "converter",
	///dentro do parenteses o nome do slice ,
	///o e a variavel com o seu ponteiro da variavel que recebera os dados
	if err != nil {
		fmt.Println("erro:", err)
	}
	fmt.Println(cesar.Nome)
	fmt.Println("fim")
}

/*
{"Nome":"cesar","Peso":62.3,"Idade":22,"Profissão":"almoxarife","Salario":2000}
*/
