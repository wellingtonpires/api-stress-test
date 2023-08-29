package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Apelido    string
	Nome       string
	Nascimento string
	Stack      []string
}

func person(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/pessoas" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "CONSULTA")
	case "POST":
		var p Person
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%+v", p)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func countPerson(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/contagem-pessoas" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "CONTAGEM")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
