package main

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

// Criando um tipo nosso
type MyNumber int

// Constraints
type Number interface {
	// Utilizamos o '~' para fazer os tipos que criarmos
	// serem aceitos pela interface
	~int | ~float64
}

// Dessa forma, conseguimos trabalhar tanto com int
// quanto com float64
// Estamos criando um tipo 'T' que aceita os dois tipos
func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// O comparable aceita os Ts desde que os Ts possam
// ser comparados
func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	// Generics

	m := map[string]int{
		"Erick": 1000,
		"Carol": 2000,
		"Ana":   3000,
	}

	m2 := map[string]float64{
		"Erick": 1000.43,
		"Carol": 2000.43,
		"Ana":   3000.43,
	}

	m3 := map[string]MyNumber{
		"Erick": 1000,
		"Carol": 2000,
		"Ana":   3000,
	}

	println(SomaInteiro(m))
	println(SomaFloat(m2))

	println(Soma(m))
	println(Soma(m2))

	println(Soma(m3))

	println(Compara(10, 10))
}
