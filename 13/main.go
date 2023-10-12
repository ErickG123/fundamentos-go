package main

import "fmt"

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
}

// Criando uma struct (tipo)
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool

	// Composição de structs
	Endereco
}

// Criando um "método" para desativar o cliente
// Devemos colocar um "nome" e a Struct antes de
// colocar um nome para a função, fazendo isso temos
// acesso as propriedades desse Struct
func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado \n", c.Nome)
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
	cliente1.Desativar()
	cliente1.Endereco.Rua = "Rua XPTO"
	cliente1.Endereco.Numero = 125
	cliente1.Endereco.Cidade = "Barra Bonita"
	cliente1.Endereco.Estado = "SP"

	fmt.Printf("Nome: %v \n", cliente1.Nome)
	fmt.Printf("Idade: %d \n", cliente1.Idade)
	fmt.Printf("Ativo: %v \n", cliente1.Ativo)

	fmt.Printf("Rua: %v \n", cliente1.Rua)
	fmt.Printf("Numero: %d \n", cliente1.Numero)
	fmt.Printf("Cidade: %v \n", cliente1.Cidade)
	fmt.Printf("Estado: %v \n", cliente1.Estado)
}
