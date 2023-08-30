package main

import (
	"fmt"
	"net/http"

	"github.com/wellingtonpires/api-stress-test/routes"
)

func main() {
	fmt.Println("Running...")
	http.HandleFunc("/pessoas", routes.PersonRoute)
	http.HandleFunc("/pessoas/", routes.PersonRoute)
	http.HandleFunc("/contagem-pessoas", routes.CountPersonRoute)
	http.ListenAndServe(":3000", nil)
}
