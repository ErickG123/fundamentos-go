package main

import "fmt"

func main() {
	// Criando um array de 3 posições do tipo inteiro
	// As posições do array são fixas, eu não consigo
	// adicionar novas casas no array
	var meuArray [3]int

	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	// Se eu tentar declrar um quarto valor, ele vai dar erro
	// falando que o excedeu o limite de tamanho do array

	// Mostrando quantas posiçõe tem o meu array, qual o tamanho dele
	fmt.Println(len(meuArray))

	// Dessa forma eu consigo mostrar a última posição do meu array
	// nesse caso, a última posição do meu array é o 2
	fmt.Println(len(meuArray) - 1)

	// Mostrando o valor na última posição do array
	fmt.Println(meuArray[len(meuArray)-1])

	// Percorrendo um array
	// Declarando o for dessa forma eu consigo pegar os
	// indices e os valores do meu array
	for i, v := range meuArray {
		// %d é utilizando quanto eu vou mostrar digitos
		// o %v eu normalmente utilizando quando é string
		fmt.Printf("O valor do indice %d é: %d \n", i, v)
	}
}
