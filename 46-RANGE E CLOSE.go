package main

import (
	"fmt"
)

///Cap. 21 – Canais – 3. Range e close
func main() {
	c := make(chan int)
	go loop(10, c)
	for v := range c {
		fmt.Println("valor do canal c:", v)
	}

}
func loop(t int, s chan<- int) { ///criando um range para poder colocar valores no canais
	for i := 0; i < t; i++ {
		s <- i
	}
	close(s) ///depois de terminar o range ,tem que utilizar o close ,
	///spara "informar" ao programa que loop acabou
}
