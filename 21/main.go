package main

import (
	"fmt"
	"pacotes/matematica"
)

func main() {
	// Eu só consigo acessar variáveis, funções e strcuts
	// se elas começarem com letra maiusculas
	resultado := matematica.Soma(10, 20)
	fmt.Printf("Resultado: %v", resultado)
	fmt.Println(matematica.A)
}
