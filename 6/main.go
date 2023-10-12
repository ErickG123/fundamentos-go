package main

import "fmt"

func main() {
	// Trabalhando com slices
	// Os slices são dividos em: ponteiros, tamanho e capacidade
	// Por baixo dos panos, os slices apontam para arrays

	// No slice eu não vou fixar um valor dentro dos colchetes
	// Eu consigo inicializar um slice vazio também
	s := []int{2, 4, 6, 8, 10}

	// len(s) mostra o tamanho do slice
	// cap(s) mostra o tamanho do slice
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)

	// Quando eu utiliza [:0] eu estou falando que a partir do inicio
	// do slice, eu quero mostrar zero items
	// O slice mantem a capacidade inicial, mesmo que eu tenho zero
	// items no slice
	fmt.Printf("len=%d cap=%d %v \n", len(s[:0]), cap(s[:0]), s[:0])

	// Neste caso, eu estou removendo todos os elementos depois da
	// quarta posição
	// Fazendo isso, meu tamanho irá para 4 e minha capacidade continuara
	// em 5
	fmt.Printf("len=%d cap=%d %v \n", len(s[:4]), cap(s[:4]), s[:4])

	// Declarando [2:] ele irá remover os 2 primeiros valores do slice
	// e ira manter o resto
	// Neste caso, a capacidade do array também irá diminuir
	// Isso acontecer porque eu ignorei o começo do array
	fmt.Printf("len=%d cap=%d %v \n", len(s[2:]), cap(s[2:]), s[2:])

	// Aumentando a capacidade do slice
	// Eu utilizo o append para adicionar um novo valor ao slice
	s = append(s, 12)
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
	// Por padrão, quando eu vou fazer um append e meu slice não tem capacidade
	// para isso, ele irá dobrar a capacidade do meu array
}
