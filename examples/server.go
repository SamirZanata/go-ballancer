package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "8001", "Porta do servidor")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Resposta do servidor: %s\n", *port)
		log.Printf("Requisição recebida no servidor %s", *port)
	})

	addr := ":" + *port
	log.Printf("Servidor iniciado na porta %s", *port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

