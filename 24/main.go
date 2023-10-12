package main

func mainc() {
	// Condicionais
	a := 1
	b := 2
	c := 3

	if a > b {
		println(a)
	} else {
		println(b)
	}

	if a > b && c > a {
		println("a > b && c > a")
	}

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	default:
		println("d")
	}
}
