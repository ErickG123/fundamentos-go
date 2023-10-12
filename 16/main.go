package main

import "fmt"

func soma(a, b int) int {
	a = 50
	return a + b
}

func somaPonteiro(a, b *int) int {
	*a = 30
	return *a + *b
}

func main() {
	num1 := 10
	num2 := 20
	// Quando fazemos isso, estamos passando uma cópia dos valores
	println(soma(num1, num2))
	println(num1)

	num3 := 20
	num4 := 40
	// Neste caso, eu estou passando o endereço na memória das variáveis
	fmt.Printf("Valor de num3 antes da soma: %d \n", num3)
	println(somaPonteiro(&num3, &num4))
	// O valor da varáivel num3 mudou pois alteramos o seu valor através
	// do seu endereço na memória (isso acontece por causa dos ponteiros)
	fmt.Printf("Valor de num3 depois da soma: %d \n", num3)
}
