package main

import (
	"fmt"
)

func main() {
	for x := 33; x < 123; x++ {
		fmt.Printf("decimal-%d,hexadecimal-%#X,unicode-%c", x, x, x)
		fmt.Println(" ")
	}

}
