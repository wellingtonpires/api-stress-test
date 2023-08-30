package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Pessoa struct {
	Apelido    string
	Nome       string
	Nascimento string
	Stack      []string
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("t")
	pathParameter := strings.Split(r.URL.Path, "/")
	if len(pathParameter) == 3 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "CONSULTA "+pathParameter[2])
	} else if t != "" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "CONSULTA"+t)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func GetCountPerson(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "CONSULTA-CONTAGEM")
}

func Create(w http.ResponseWriter, r *http.Request) {
	var p Pessoa
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%+v", p)
}
