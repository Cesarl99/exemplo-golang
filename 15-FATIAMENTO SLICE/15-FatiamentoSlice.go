package main

import (
	"fmt"
)

///fatiamento de slice
func main() {

	materiais := []string{"lapis", "caneta", "lousa", "borracha", "caderno"}
	compra := materiais[0:2]
	///esse exemplo queremos recortar o item na posição 0(lapis) e o 1(caneta),os demais itens
	///o item selecionado no final não é recortado somente ate a posição antes
	ncomprar := materiais[2:]
	///nesse exemplo comeca no 2 posição e para ir ate o final da slice deixamos em branco
	fmt.Println(compra)
	fmt.Println(ncomprar)

	for i := 0; i < len(materiais); i++ {
		fmt.Println(materiais[i])
	}
}
