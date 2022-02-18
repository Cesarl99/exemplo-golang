package main

import (
	"fmt"
)

///slice
func main() {
	slice := []int{1, 3, 5, 6, 1, 8}
	total := 0
	fmt.Println(slice)

	slice2 := append(slice, 99)
	slice2 = append(slice2, 666)
	slice2 = append(slice2, 696)
	///é utilizada na criação de da slice e acrecenta com append ,
	/// neste caso acrecentando o slice em outro slice
	for indice, valor := range slice2 {
		///o range "corre" todo o slice e executa a função a baixo
		fmt.Println("no indice", indice, "temos o valor:", valor)
		total += valor
		/// o total soma do o valor da slice cada vez que ele passa pelo range
	}

	fmt.Printf("o soma de todos os valores da slice é :%v", total)
}
