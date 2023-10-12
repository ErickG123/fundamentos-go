package main

import "fmt"

func main() {
	// Funções variáticas
	fmt.Println(sum(10, 20, 40, 60))
}

// Utilizo os 3 pontinhos quando eu não sei
// quantos parametros minha função vai receber
// Nesse caso, ela aceita uma infinidade de
// parametros, desde que eles sejam inteiros
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
