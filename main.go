package main

import (
	"crud-go/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", server.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", server.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", server.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", server.AtualizarUsuarios).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", server.DeletarUsuario).Methods(http.MethodDelete)

	fmt.Println("Listen Port: 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
