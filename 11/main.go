package main

import "fmt"

// Criando uma struct (tipo)
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	// Structs
	// Uma Struct não é uma Classe, mas se assemelha a uma
	cliente1 := Cliente{
		Nome:  "Erick",
		Idade: 20,
		Ativo: true,
	}

	// Mudando o valor de uma struct
	cliente1.Ativo = false

	fmt.Printf("Nome: %v \n", cliente1.Nome)
	fmt.Printf("Idade: %d \n", cliente1.Idade)
	fmt.Printf("Ativo: %v \n", cliente1.Ativo)
}
