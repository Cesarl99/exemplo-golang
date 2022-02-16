package main

import (
	"fmt"
)

func main() {
	x := []int{1, 5, 6, 3, 8, 1}
	y := []int{5, 5, 100, 1, 777}
	fmt.Println(x)
	fmt.Println(y)
	x = append(x, y...)
	///para podemos juntar uma slice dentro de outra utilizamos ,nome da slice com ...
	fmt.Println(x)

}
