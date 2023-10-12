package main

// Utilizando o pacote fmt para trabalhar com formatação
import "fmt"

type ID int

var (
	ident ID     = 1
	nome  string = "Erick"
	idade int    = 20
)

func main() {
	// O "%T" traz o tipo das variáveis
	// O "%v" traz o valor das variáveis
	fmt.Printf("O tipo de ident é: %T \n", ident)
	fmt.Printf("O tipo de ident é: %v \n", ident)

	fmt.Printf("O tipo de nome é: %T \n", nome)
	fmt.Printf("O tipo de nome é: %v \n", nome)

	fmt.Printf("O tipo de idade é: %T \n", idade)
	fmt.Printf("O tipo de idade é: %v \n", idade)
}
