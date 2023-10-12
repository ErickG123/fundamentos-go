// Mesmo nome do diretório
package matematica

func Soma[T int | float64](a, b T) T {
	return a + b
}

var A int = 10
var b int = 20

type Carro struct {
	Marca string
}
