package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

///Cap. 21 – Canais – 7. Divergência
///e ao contrario de convergencia ,divergencia e quando temos um canal e queremos separar
func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)

	go manda(20, canal1)
	go outra(canal1, canal2)

	for v := range canal2 {
		fmt.Println(v)
	}

}

func manda(n int, canal chan int) { ///1-alimetamos o primeiro canal
	for i := 0; i < n; i++ {
		canal <- i
	}
	close(canal)
}

func outra(canal1, canal2 chan int) {
	var wg sync.WaitGroup
	for v := range canal1 { ///2-utilizamos os valores na canal 1,e "damos um GOROUTINES PARA CADA", sendo assim ele alimenta
		wg.Add(1) /// o trabalho/canal2(linha 38)
		go func(x int) {
			canal2 <- trabalho(x)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(canal2)
}

func trabalho(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return n
}

///exemplo usamos a digencia para separar os valores do canal,e depois a convergencia para o juntar no canal 2
