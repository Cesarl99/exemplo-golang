package main

import (
	"fmt"
)

///structs anonimos
func main() {
	x := struct {
		nome  string
		idade int
		peso  float64
	}{
		nome:  "cesar",
		idade: 22,
		peso:  66.5}

	fmt.Println(x)
}
