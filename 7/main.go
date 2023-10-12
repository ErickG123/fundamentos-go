package main

import "fmt"

func main() {
	// Trabalhando com maps (chave/valor)
	// No map eu declrado o tipo da chava e depois o tipo do valor
	// map[tipo_chave]tipo_valor{chaves/valores}
	salarios := map[string]int{"Erick": 3000, "Carol": 10000, "Ana": 5000}

	// Imprimindo o valor da chave
	fmt.Println(salarios["Erick"])

	// Deletando um valor pela chave
	delete(salarios, "Ana")
	fmt.Println(salarios)

	// Adicionando valores ao map
	salarios["Marcio"] = 20000
	fmt.Println(salarios)

	// Utilizando a função make
	// Utilzamos o make para inicializar o map vazio
	salarios_make := make(map[string]int)
	// Ou...
	salarios_vazio := map[string]int{}

	fmt.Println(salarios_make)
	fmt.Println(salarios_vazio)

	// Percorrendo um map
	for nome, salario := range salarios {
		fmt.Printf("O salario de %v é: %d \n", nome, salario)
	}

	// Ignorando um valor do map, nesse caso a chave
	// O "_" é um blank identifier
	for _, salario := range salarios {
		fmt.Printf("O salario é de: %d \n", salario)
	}
}
