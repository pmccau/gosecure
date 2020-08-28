package main

import (
	"fmt"
	"gosecure/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on port 8080...")

	port := ":8080"
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
