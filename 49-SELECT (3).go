package main

import "fmt"

///Cap. 21 – Canais – 4. Select

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
	close(impar)
	close(par)
	quit <- true
}
func raceive(par, impar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("o numero", v, "é par.")
		case v := <-impar:
			fmt.Println("o numero", v, "é impar.")
		case <-quit:
			return

		}
	}
}
