package main

import (
	"fmt"
	"runtime"
	"sync"
)

///Cap. 18 – Concorrência – 4. Na prática: Condição de corrida
///condição de corrida e quando esse uma concorrencia e as função que fazem parte dessa concorrencia
///recebem da mesma variavel de maneira desordenada podendo apresentar erro
///para verificar se há uma condição de corrida e so digitar no terminal -race antes do nome do programa
///MUTEX
///PARA RESOLVER ESSE PROBLEMA UTILIZAMOS O MUTEX.a sua funcionalidade é "trancar" uma parte de codigo
///para apenas uma goroutines  rode por vez
func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	contador := 0
	totaldegoroutines := 1000

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	var mx sync.Mutex

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			mx.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mx.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)

}
