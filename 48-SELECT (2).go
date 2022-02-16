package main

import (
	"fmt"
)

///Cap. 21 – Canais – 4. Select
///esse exemplo,mostra que ele serve tanto enviar quando receber informações
func main() {
	canal := make(chan int)
	quit := make(chan int)
	go recebequit(canal, quit)
	enviarPracanal(canal, quit)
}
func recebequit(canal chan int, quit chan int) {
	for i := 0; i < 50; i++ {
		fmt.Println("recebido", <-canal)
	}
	quit <- 0
}
func enviarPracanal(canal chan int, quit chan int) {
	qualquercoisa := 1
	for {
		select {
		case canal <- qualquercoisa:
			qualquercoisa++
		case <-quit:
			return
		}
		fmt.Println("qualquircoisa", qualquercoisa)
	}
}
