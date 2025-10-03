package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gustavolubacheski/API/CRUD-API/routes"
)

func main() {
	fmt.Println("Listening on port 8080")

	r := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", r)) // Single blocking call with your router
}
