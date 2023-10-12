package main

import "fmt"

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
}

// Utilizando uma Interface
// No Go, só conseguimos passar métodos nas Interfaces
// não é possível passar propriedades
type Pessoa interface {
	// Fazendo isso eu falo que qualquer Struct que
	// tenha o método Desativar() está implementando
	// a interface de Pessoa
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado \n", c.Nome)
}

// Criando uma função que vai receber a Interface como
// o tipo do parâmetro
// Garante que qualquer Struct que tenha o método desativar
// possa ser utilizado
func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	cliente1 := Cliente{
		Nome:  "Erick",
		Idade: 20,
		Ativo: true,
	}

	// Eu consigo fazer isso, pois a Struct de Cliente está
	// recebendo a implementação da Interface de Pessoa de
	// forma automática
	Desativacao(cliente1)
}
