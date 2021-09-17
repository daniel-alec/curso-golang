package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 12344.32,
			"Guga Pereira":   83223.0,
		},
		"J": {
			"José João": 1123.45,
		},
		"P": {
			"Pedro Junior": 1200.0,
		},
	}

	delete(funcPorLetra, "P")

	for letra, funcs := range funcPorLetra {
		fmt.Println(letra, funcs)
	}

	fmt.Println("-------")
	for letra, funcs := range funcPorLetra {
		for nome, valor := range funcs {
			fmt.Println(letra, nome, valor)
		}
	}
}
