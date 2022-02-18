package main

import (
	"fmt"
)

///Cap. 12 – Funções – 5. Interfaces & polimorfismo
type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}
type jogadorfutebol struct {
	pessoa
	timequejoga string
	gols        int
}
type jogadorbasquete struct {
	pessoa
	posição     string
	totalcestas int
}

func (x jogadorfutebol) apresentacao() {
	fmt.Println("meu nome é :", x.nome, "\n e eu jogo no ", x.timequejoga, "\n e fiz ", x.gols, "gols ")
}
func (x jogadorbasquete) apresentacao() {
	fmt.Println("meu nome é :", x.nome, "\n e eu jogo no ", x.posição, "\n e fiz ", x.totalcestas, "cestas")
}

type gente interface {
	oibomdia()
}

func atletas(x gente) {

	switch x.(type) {
	case jogadorfutebol:
		fmt.Println("fiz ", x.(jogadorfutebol).gols, "gols")

	case jogadorbasquete:
		fmt.Println("fiz ", x.(jogadorbasquete).totalcestas, "cestas")
	}
	x.oibomdia()
}

func main() {
	pele := jogadorfutebol{
		pessoa: pessoa{
			nome:      "pele",
			sobrenome: "pele",
			idade:     80,
		},
		timequejoga: "santos",
		gols:        1300,
	}
	lebron := jogadorbasquete{
		pessoa: pessoa{
			nome:      "lebron",
			sobrenome: "james",
			idade:     30,
		},
		posição:     "ala",
		totalcestas: 3600,
	}
	pele.apresentacao()
	lebron.apresentacao()
	atletas(lebron)
	atletas(pele)
}
