package main

import (
	"fmt"
)

///resolvendo o problema do exemplo 53 ,com Cap. 21 – Canais – 5. A expressão comma ok

func main() {
	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)
	go mandarNumeros(par, impar, quit)
	raceive(par, impar, quit)
}
func mandarNumeros(par, impar chan int, quit chan bool) {
	total := 200
	for i := 0; i < total; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
	quit <- true
}
func raceive(par, impar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("o numero", v, "é par.")
		case v := <-impar:
			fmt.Println("o numero", v, "é impar.")
		case v, ok := <-quit:
			if !ok {
				fmt.Println("Deu zebra! Ó:", v)
			} else {
				fmt.Println("Encerrando. Recebemos:", v)
			}
			return
		}
	}
}

/*para resolver o erro de apareceu mais numeros do que deveria utilizamos o expressão comma ok(linha 36),
que verifica se foi recebido no ok,se for diferente do correto(linha 37),print erro,já se for igual a correto(linha39),print encerrando
*/
