package main

import (
	"fmt"
)

///maps
func main() {
	amigos := map[string]int{
		"cesar":   36363651,
		"luiz":    145874586,
		"joaquim": 25631457,
	}
	fmt.Println(amigos)
	fmt.Println(amigos["jose"])
	for key, value := range amigos {
		fmt.Println(key, value)
	}
	delete(amigos, "luiz")
	///utilize o delete para apagar um key de um map,
	fmt.Println(amigos)
}
