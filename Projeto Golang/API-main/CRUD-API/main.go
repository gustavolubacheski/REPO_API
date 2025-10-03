package main

import (
	"fmt"
	"net/http"

	"github.com/gustavolubacheski/API/CRUD-API/routes"
)

func main() {
	fmt.Println("Listening on port 8080")

	r := routes.NewRouter()

	http.ListenAndServe(":8080", r)
}
