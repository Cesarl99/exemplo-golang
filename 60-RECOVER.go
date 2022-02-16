package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

///Cap. 23 – Tratamento de Erros – 4. Recover
func f() {
	defer func() { ///<------------------------ultima a ser executada
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("tirando o g", i) //<---------primeira a ser executada
	fmt.Println("mostrando o g", i)
	g(i + 1)
}
