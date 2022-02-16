package main

import (
	"fmt"
)
///metodos
type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) cadastro() {
	fmt.Println(p.nome, "bom dia")
}
func main() {
	cesar := pessoa{"cesar", 30}
	cesar.cadastro()

}
