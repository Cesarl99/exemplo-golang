package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

/// Cap. 16 – Aplicações – 7. bcrypt

func main() {
	senha := "01janeiro2000"

	sb, err := bcrypt.GenerateFromPassaword
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sb))
}
