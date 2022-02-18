package main

import (
	"fmt"
	"sort"
)

///Cap. 16 – Aplicações – 5. O pacote sort
/*o pacote sort e importado e ter diversas funções (de ordenação) no exemplo abaixo utilizamos o strings
(que classifica em ordem crescente as variaveis função strings)
e a função int (que classifica em ordem crescente as variaveis tipo int)
*/
func main() {
	ss := []string{"bola", "campo", "camisa", "time", "musica", "copa"}
	si := []int{2, 6, 2, 1, 0, 5, -1, 3, 5, 1, 4, 100, 214}
	sort.Strings(ss)
	sort.Ints(si)
	fmt.Println(ss)
	fmt.Println(si)
}
