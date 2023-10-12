package main

import "fmt"

func main() {
	// Trabalhando com funções anonimas (closures)
	// Dentro de uma função podemos ter outra função

	// Normalmente utilizamos funções anonimas no Go para
	// executar rotinas, por exemplo, rodar um web server

	// Não damos um nome a função, só vamos colocar seu tipo
	total := func() int {
		return sum(10, 20, 40, 60) * 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
