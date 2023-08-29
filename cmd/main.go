package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Running...")
	http.HandleFunc("/pessoas", person)
	http.HandleFunc("/contagem-pessoas", countPerson)
	http.ListenAndServe(":3000", nil)
}
