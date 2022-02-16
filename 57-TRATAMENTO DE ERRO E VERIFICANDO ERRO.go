/*
package main

import "fmt"

////Cap. 23 – Tratamento de Erros – 2. Verificando erros
quando se existe um erro ,há dois caminhos ,se o erro não existir o programa executa normalmente
quando há um erro,podemos tratar esse mostrando esse erro ou dando panic

func main() {
	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}

package main

import "fmt"

func main() {
	var answer1, answer2, answer3 string

	fmt.Print("Name: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Sport: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer1, answer2, answer3)

}*/