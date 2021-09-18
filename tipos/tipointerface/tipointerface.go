package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	// inteface vazia: pode atribuir tipos genéricos
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	// ou tipo interface vazia
	type dinamico interface{}

	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Golang: Explorando a Linguagem do Goole"}
	fmt.Println(coisa2)
}
