package main

func main() {
	// Looping
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []int{1, 2, 3, 4, 5}
	for k, v := range numeros {
		println(k, v)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	for {
		println("Erick")
	}
}
