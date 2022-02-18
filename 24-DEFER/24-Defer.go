package main

import (
	"fmt"
)

///defer
func main() {
	defer fmt.Println("primeiro")
	fmt.Println("segundo")
	defer fmt.Println("terceiro")
	fmt.Println("quarto")
}
