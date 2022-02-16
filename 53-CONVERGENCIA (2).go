package main

import (
	"fmt"
	"math/rand"
	"time"
)

////Cap. 21 – Canais – 6. Convergência

func main() {
	canal := converge(trabalho("maçã"), trabalho("pêra"))
	for x := 0; x < 16; x++ {
		fmt.Println(<-canal)
	}

}

func trabalho(s string) chan string {
	canal := make(chan string)
	go func(s string, c chan string) {
		for i := 1; ; i++ {
			c <- fmt.Sprintf("Função %v diz: %v", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(s, canal)
	return canal
}

func converge(x, y chan string) chan string {
	novo := make(chan string)
	go func() {
		for {
			novo <- <-x
		}
	}()
	go func() {
		for {
			novo <- <-y
		}
	}()
	return novo
}

/*
esse exemplo se trata de uma convergencia por que o inicio se da apartir de duas func ambas
chamadas de trabalho ,com canais chamados de maçã e pera,
e apartir dela elas se juntam em uma chamada de converge



*/
