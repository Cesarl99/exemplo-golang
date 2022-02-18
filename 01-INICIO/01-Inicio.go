package main

import (
	"fmt"
)

var soma int
var div float32

func main() {
	v := 22
	calculosoma(v)
}
func calculosoma(x int) {
	soma = x + 10
	fmt.Println("a \"soma\" é: ", soma)
}
