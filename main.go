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
	garagePin.Detect(rpio.RiseEdge)
	frontDoorPin := rpio.Pin(23)
	frontDoorPin.Input()
	frontDoorPin.Detect(rpio.RiseEdge)

	for {
		if frontDoorPin.EdgeDetected() {
			fmt.Println("Edge at front door!")
		}
	}

//	for {
//		garage := garagePin.Read()
//		frontDoor := frontDoorPin.Read()
//		if frontDoor == 1 {
//			fmt.Println("Exiting, front door hot!")
//			break
//		}
//		if garage == 1 {
//			fmt.Println()
//		}
////		fmt.Println("Garage res:", garagePin.Read())
////		fmt.Println("Front res:", frontDoorPin.Read())
//	}

	r := router.Router()
	fmt.Println("Starting server on port 8080...")

	port := ":8080"
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
