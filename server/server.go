package server

import (
	"crud-go/database"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID      uint32 `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Created string `json:created`
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := database.InitDB()
	if erro != nil {
		log.Println("Erro ao estabelecer conexão com o banco de dados:", erro)
		w.Write([]byte("Erro ao estabelecer conexão com o banco de dados"))
		return
	}
	defer db.Close()

	query := "select * from users"
	linhas, err := db.Query(query)
	if err != nil {
		log.Println("Erro ao executar query:", err)
		w.Write([]byte("Erro ao executar query"))
		return
	}

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Name, &usuario.Email, &usuario.Created); erro != nil {
			w.Write([]byte("Erro ao escanear usuario"))
			log.Println("Erro ao executar query:", erro)
			return
		}
		usuarios = append(usuarios, usuario)
	}
	w.WriteHeader(http.StatusOK)
	response := json.NewEncoder(w).Encode(usuarios)
	if response != nil {
		w.Write([]byte("Erro ao converter os usuarios para o JSON"))
		return
	}

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter parametro"))
	}

	db, erro := database.InitDB()
	if erro != nil {
		log.Println("Erro ao estabelecer conexão com o banco de dados:", erro)
		w.Write([]byte("Erro ao estabelecer conexão com o banco de dados"))
		return
	}
	defer db.Close()

	linha, erro := db.Query("select * from users where id = $1", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuario"))

	}
	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Name, &usuario.Email, &usuario.Created); erro != nil {
			w.Write([]byte("Erro ao escanear usuario"))
			log.Println("Erro ao executar query:", erro)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	response := json.NewEncoder(w).Encode(usuario)
	if response != nil {
		w.Write([]byte("Erro ao converter os usuarios para o JSON"))
		return
	}
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario
	erro = json.Unmarshal(corpoRequisicao, &usuario)
	if erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	/* forma mais elegante de escrever o mesmo bloco de código acima em go

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
	}

	*/

	db, erro := database.InitDB()
	if erro != nil {
		log.Println("Erro ao estabelecer conexão com o banco de dados:", erro)
		w.Write([]byte("Erro ao estabelecer conexão com o banco de dados"))
		return
	}
	defer db.Close()
	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	var idInserido uint32
	err := db.QueryRow(query, usuario.Name, usuario.Email).Scan(&idInserido)
	if err != nil {
		log.Println("Erro ao executar o statement:", err)
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso: Id: %d", idInserido)))

}
