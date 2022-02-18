package main

import (
	"fmt"
)

///Cap. 21 – Canais – 5. A expressão comma ok
///a expressão comma ok e utilizado para verificar se o valor recebido da canal e realmente um valor verdadeiro
///

func main() {
	canal := make(chan int)

	go func() {
		canal <- 42
		close(canal)
	}()
	v, ok := <-canal ///verificando o valor que esta no canal e se ele e verdadeiro ou não
	///nesse caso e verdadeiro porque e está retirando o valor do canal e sendo mostrado
	fmt.Println("o numero que esta no canal e ", v, "e o valor é :", ok)

	v, ok = <-canal ///nesse caso o valor que esta canal ja não esta mais por isso
	/// o valor mostardo sera falso

	fmt.Println("o numero que esta no canal e ", v, "e o valor é :", ok)

}
