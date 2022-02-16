package main

import (
	"fmt"
	"sync"
)

////Cap. 21 – Canais – 6. Convergência
///convergencia e quando temos duas canais diferentes e unimos em somente um
func main() {
	par := make(chan int)
	converge := make(chan int)
	impar := make(chan int)

	go envia(par, impar)
	go receber(par, impar, converge)

	for v := range converge {
		fmt.Println(v)
	}
}
func envia(p, i chan int) {
	x := 100
	for y := 0; y < x; y++ {
		if y%2 == 0 {
			p <- y
		} else {
			i <- y
		}
	}
	close(i)
	close(p)
}
func receber(p, i, c chan int) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for v := range p {
			c <- v
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for v := range i {
			c <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(c)
}
