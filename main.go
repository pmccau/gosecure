package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"gosecure/router"
	"log"
	"net/http"
)

func main() {

	err := rpio.Open()
	if err != nil {
		log.Fatal("Error opening pins", err)
	}
	garagePin := rpio.Pin(18)
	garagePin.Input()
	frontDoorPin := rpio.Pin(23)
	frontDoorPin.Input()

	for {
		fmt.Println("Garage res:", garagePin.Read())
		fmt.Println("Front res:", frontDoorPin.Read())
	}

	r := router.Router()
	fmt.Println("Starting server on port 8080...")

	port := ":8080"
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
