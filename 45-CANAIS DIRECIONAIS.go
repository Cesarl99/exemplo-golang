package main

import (
	"fmt"
)

///Cap. 21 – Canais – 2. Canais direcionais & Utilizando canais
///canais podem ser direcionais ,ou seja ,pode ser somente para retirar informações ou so para receber informções
func main() {
	canal := make(chan int)
	go send(canal) ///para rodar varios canais ,não podemos rodar uma de cada vez ,elas devem rodar de maneira concorrente(gorotines),ou
	//seja definir ao menos um canal como go
	receive(canal)
}

func send(s chan<- int) { ///funcão criado para receber informação(send)
	s <- 42
}
func receive(r <-chan int) { ///funcão criado para retirar informação(receive)
	fmt.Println("o valor recebido do canal foi:", <-r)
}
