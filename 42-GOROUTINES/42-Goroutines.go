package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
gorotines é o comando dado para executar varias funções ao mesmo tempo
 Um WaitGroup serve para esperar que uma coleção de goroutines termine sua execução.
    - func Add: "Quantas goroutines?"
    - func Done: "Deu!"
    - func Wait: "Espera todo mundo terminar."
- Ah, mas então... sim!
- Só pra ver: runtime.NumCPU() & runtime.NumGoroutine()
*/
//Para Criar uma goroutines seguimos o seguintes passsos:
var wg sync.WaitGroup //1-criamos uma variavel

func main() {

	fmt.Println(runtime.NumCPU())       //utilizado para verificar numero de processadores que estão sendo utilizados
	fmt.Println(runtime.NumGoroutine()) //numero de gorotines ate aquele momento

	wg.Add(2) //2-"Mostramos" para o programa quantos goroutines existem

	go func1() //3-determinamos quais funções são goroutines
	go func2()

	fmt.Println(runtime.NumGoroutine()) //numero de gorotines ate aquele momento

	wg.Wait() //4-um aviso para esperar todos o goroutines rodar

}

func func1() {
	for i := 0; i < 100; i++ {
		fmt.Println("func1:", i)
		time.Sleep(20)
	}
	wg.Done() //5-adcionamos essa função a waitgroup
}

func func2() {
	for i := 0; i < 100; i++ {
		fmt.Println("func2:", i)
		time.Sleep(20)
	}
	wg.Done() //5-adcionamos essa função a waitgroup
}
