package main

import (
	"fmt"
	"sync"
)

var mx sync.Mutex

///Cap. 21 – Canais – 4. Select
///o select tem a função parecida ao case
func main() {
	a := make(chan int)
	b := make(chan int)
	x := 100

	go func(x int) {
		for i := 0; i < x; i++ {
			a <- i
		}
	}(x / 2)

	go func(x int) {
		for i := 0; i < x; i++ {
			b <- i
		}
	}(x / 2)

	for i := 0; i < x; i++ { ///o select permite realizar uma ação diferente para cada ação
		///   que acontecer
		select {
		case v := <-a: ///nesse caso se o valor esta no canal A ,vai escrever  "CANAL A"+ VALOR
			fmt.Println("canal A:", v)
		case v := <-b: ///nesse caso se o valor esta no canal B ,vai escrever  "CANAL B"+ VALOR
			fmt.Println("canal B:", v)
		}
	}
	///desse jeito os valores colocados na tela nos dois case aparece de maneira aleatoria devido ao go
}
