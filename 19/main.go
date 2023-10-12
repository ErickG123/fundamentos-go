package main

import "fmt"

func main() {
	// Type Assertations
	// Neste caso, a interface não definite o tipo da var
	var minhaVar interface{} = "Erick"

	// Fazendo isso, eu estou falando para o programa
	// qual é de fato o tipo da variável
	println(minhaVar.(string))

	// Neste caso, eu estou pegando o resultado da conversão
	// e var 'ok' vai armazenar true se a conversão der certo
	// e false a conversão der errado
	resultado, ok := minhaVar.(int)
	fmt.Printf("O resultado é: %v \nO valor de ok é: %v \n", resultado, ok)

	// Neste caso, como eu não tenho o 'ok' para verificar se a
	// conversão deu certo, o código vai dar erro
	resultado2 := minhaVar.(int)
	fmt.Printf("O resultado2 é: %v \n", resultado2)
}
