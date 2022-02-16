package main

import (
	"fmt"
	"log"
	"os"
)

/*
Cap. 23 – Tratamento de Erros – 3. Print & Log
exemplos de println,log.Println,log.Setoutput,log.Panicln
*/
func main() {
	f, err := os.Create("logpanic.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f) //<------------------------------recebe um writer(qualquer interface writer),esse caso e "diz" que
	//todo o erro que pode dar ele salva nesse arquivo

	f2, err := os.Open("no-file.txt ")
	if err != nil {
		//		fmt.Println("err happened", err)//<---printa somente o erro
		//		log.Println("houve um erro aqui", err) //<---Printa o erro com a hora e a data do erro
		//		log.Fatalln(err)
		//		log.Panicln(err)//<-----------------------------util para printar mais detalhadamente um erro
		//		panic(err)//<-----------------------------a função panic ,ele encerra a função que era esta rodando se tiver erro
		//mas as demais função consegue rodar normalmente
	}
	defer f2.Close()

	fmt.Println("check the log.txt file in the directory")
}
