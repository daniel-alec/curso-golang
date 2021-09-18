package main

import "fmt"

// uma função também é um closure
func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x) // lê o 'x' do scopo deste função
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX()
}
