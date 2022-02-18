package main

import (
	"fmt"
	"sort"
)

///Cap. 16 – Aplicações – 6. Customizando o sort
///para custumizar o sort temos:
///1- que criar uma struct
type jogador struct {
	nome   string
	gols   int
	titulo int
}

///2-criar um type para ordenar em as informação da struct
type ordenarPorgols []jogador

func (x ordenarPorgols) Len() int           { return len(x) }
func (x ordenarPorgols) Less(i, j int) bool { return x[i].gols > x[j].gols }
func (a ordenarPorgols) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

///2-criar um type para ordenar em as informação da struct
type ordenarPortitulo []jogador

func (x ordenarPortitulo) Len() int           { return len(x) }
func (x ordenarPortitulo) Less(i, j int) bool { return x[i].titulo > x[j].titulo }
func (a ordenarPortitulo) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

///3-na função main temos que entra com os dados da struct que a torna um slice
func main() {
	jogadores := []jogador{jogador{"naymar", 610, 8},
		jogador{"pele", 1200, 15},
		jogador{"jão", 36, 16},
	}
	fmt.Println(jogadores)
	///4-chamamos a função sort que a sua customização e chamada de sort mas com a type pre definido
	sort.Sort(ordenarPorgols(jogadores))

	fmt.Println("os jogadores em ordem de que fez mais gols e :\n", jogadores)
	fmt.Printf("--------------------------------------------------------------------------\n")
	///4-chamamos a função sort que a sua customização e chamada de sort mas com a type pre definido
	sort.Sort(ordenarPortitulo(jogadores))
	fmt.Println("os jogadores em ordem de que fez mais titulos e :\n", jogadores)

}
