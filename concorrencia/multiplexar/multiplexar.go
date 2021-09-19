package main

import (
	"fmt"

	"github.com/daniel-alec/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem // dados do canal origem para canal destino
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	// um canal para juntar outros canais
	c := juntar(
		html.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
		html.Titulo("https://www.amazon.com", "https://www.youtube.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
