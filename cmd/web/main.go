package main

import (
	"fmt"
	"go-web/pkg/handlers"
	"net/http"
)

const PORT = ":8080"




func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Listening on port %s", PORT)

	http.ListenAndServe(PORT, nil)
}