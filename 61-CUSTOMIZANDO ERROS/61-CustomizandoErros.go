package main

import (
	"errors"
	"fmt"
	"log"
)

///Cap. 23 – Tratamento de Erros – 5. Erros com informações adicionais
///customizando erros
var errosnew = errors.New("não existe raiz quadrada de numero negativo") ///<------com o errors.new,podemos costumizar a mensagem quando há um erro

func main() {
	fmt.Printf("%t\n", errosnew)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

}
func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errosnew /// <---------e utilizamos a mensagem do erro que foi colocado em uma variavel em qualquer lugar do programa
	}
	return 42, nil
}
