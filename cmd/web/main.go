package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":9090"

type application struct {
}

func main() {
	// app := &application{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello "))
	})

	fmt.Println("Starting server on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Panic(err)
	}
}
