package main

import "fmt"

type Cliente struct {
	nome string
}

func (c Cliente) andou() {
	c.nome = "Erick Gabriel"
	fmt.Printf("O Cliente %v andou \n", c.nome)
}

type Conta struct {
	saldo int
}

// Passando o endereço de onde a conta está armazenada
func (ct *Conta) simular(valor int) int {
	ct.saldo += valor
	fmt.Printf("O valor da silumação é: %d \n", ct.saldo)
	return ct.saldo
}

// Retornamos o Endereço da nossa Struct, para que as alterações
// que fizermos nela sejam permanentes
func NovaConta() *Conta {
	return &Conta{
		saldo: 0,
	}
}

func main() {
	// Structs e Ponteiros
	cliente1 := Cliente{
		nome: "Erick",
	}
	cliente1.andou()

	fmt.Printf("O valor da struct com nome %v \n", cliente1.nome)

	conta1 := Conta{
		saldo: 100,
	}
	conta1.simular(200)
	fmt.Printf("Valor pós simulação: %d \n", conta1.saldo)
}
