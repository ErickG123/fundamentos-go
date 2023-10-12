package main

import "fmt"

func main() {
	// Memória -> Endereço -> Valor
	// Variável -> Ponteiro -> Endereço -> Valor
	a := 10
	fmt.Println(a)

	// Vendo o endereço da memória da variável
	fmt.Println(&a)

	// Quanto utilizamos o '*' estamos falando do endereço
	// de memória da variável
	var ponteiro *int = &a
	fmt.Println(ponteiro)

	// Atribuindo um novo valor ao endereço de memória
	*ponteiro = 20
	fmt.Println(a)

	// Neste caso, 'b' está apontando para o mesmo endereço
	// de 'a'. Está funcionando como um ponteiro
	b := &a
	fmt.Println(b)

	// Verificando o valor de 'b'
	// Neste caso, o '*' é utilizado para acessar o valor
	fmt.Println(*b)

	*b = 30
	fmt.Println(*b)
	fmt.Println(a)
}
