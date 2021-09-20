package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05") // cada um tem um significado  02->mes, 01->dia, 03->hora, 04->min, 05->seg, 06->ano + 20
	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", s)
}

func main() {
	http.HandleFunc("/", horaCerta) // lógica para consumir conteúdos dinâmicos
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
