package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sub(1, 2))

	// Colocando os retornos da função em variáveis separadas
	// Esse tipo de tratamento é muito comum no Go
	divisao, erro := div(10, 0)
	if erro != nil {
		fmt.Println(erro)
	}
	fmt.Println(divisao)
}

// Criando funções
// Preciso especificar os tipos do parametros
// e o tipo de retorno da função, se eu deixar
// o retorno vazio, quer dizer que ela não retorna
// nada
func sum(a int, b int) int {
	return a + b
}

// Podemos criar uma função que retorna que retorna mais
// de um valor
// Para passar mais de um tipo eu utilizo os parenteses
func sub(a int, b int) (int, bool) {
	if a-b >= 50 {
		return a - b, true
	} else {
		return a - b, false
	}
}

// Normalmente eu uso dois retornos para tratar erro
// Para isso eu passo o segundo tipo como "error"
func div(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Não existe divisão por zero")
	} else {
		return a / b, nil
	}
}
