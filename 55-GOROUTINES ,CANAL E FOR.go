package main

import (
	"fmt"
	"sync"
	"time"
)

///nesse exemplo e util quando não queremos rodar muitas GOROUTINES  de uma so vez
func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)
	funções := 5
	go manda(100, canal1)
	go outra(funções, canal1, canal2)
	for v := range canal2 {
		fmt.Println(v)
	}
}

func manda(n int, canal chan int) {
	for i := 0; i < n; i++ {
		canal <- i
	}
	close(canal)
}

func outra(funções int, canal1, canal2 chan int) { ///definimos o valor da função,ou seja quantas goroutines queremos rodar de uma so vez
	var wg sync.WaitGroup ///ela rodar a quantidade definida e quando termina e roda novamente essa quantidade,ate atingir o valor de n digitado na func main
	for i := 0; i < funções; i++ {
		wg.Add(1)
		go func() {
			for v := range canal1 {
				canal2 <- trabalho(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(canal2)
}

func trabalho(n int) int {
	time.Sleep(time.Millisecond * 1000) //time.Duration(rand.Intn(1e3)))
	return n
}
