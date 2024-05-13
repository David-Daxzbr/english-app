package main

import (
	"fmt"
	"net/http"

	"github.com/david-daxzbr/english-app/handlers"
)

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/login", handlers.Login)

	fmt.Println("Listening on port :3000")
	http.ListenAndServe(":3000", nil)
}
