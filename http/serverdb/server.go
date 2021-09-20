package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/usuarios/", UsuarioHandler) // vai achar rodando na linha de comando
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// cd serverdb
// go run *.go
