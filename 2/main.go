package main

// Declaração de uma constante, o valor
// não pode ser mudado ao longo do código
const a = "Hello, World"

// Declarando várias variáveis com o parenteses
// Essas variáveis são de escopo global, todas
// as funções tem acesso a elas

// Variáveis declarando com "var" podem ter o valor
// alterado
var (
	b bool    // A declaração de bool por padrão vem como false
	c int     // A declaração de int por padrão vem como 0
	d string  // A declaração de string por padrão vem em branco
	e float64 // A declaração de float64 por padrão vem como 0.00

	f bool = true // Declarando tipo e atribuindo valor a variável
)

func main() {
	// Atribuindo um valor a variável
	b = true
	c = 100
	d = "Oi"
	e = 1.78

	// Variável de escopo local
	var x string

	// Simplicando a criação de variáveis
	// O ":=" é usado só uma vez na variável, que é quando
	// ela está sendo criado, depois disso eu só uso o sinal
	// de atribuição (=)
	z := 42 // De forma ele entende o tipo automaticamente

	// Imprimindo o valor
	println(b)
	println(c)
	println(d)
	println(e)

	println(x)

	println(z)
}
