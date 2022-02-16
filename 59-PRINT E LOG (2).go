package main

import (
	"fmt"
	"log"
	"os"
)

/*
Cap. 23 – Tratamento de Erros – 3. Print & Log
exemplos de LOG.fatalln
*/

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		log.Fatalln(err) //<-------------------------quando usasse o fatalln ele encerra o programa de
		//maneira imediada,e o programa da um resultado diferente de 1 ,quando há um erro
		//		panic(err)
	}
}

func foo() {
	fmt.Println("When os.Exit() is called, deferred functions don't run")
}
