package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"log"
	"net/http"
)

type Pin struct {
	Name 			string
	Number 			int
	AtRest			bool
	Current			bool
	Pin 			rpio.Pin
}

var gpioInit = false
var Pins []*Pin

// NewPin is the constructor for a Pin
func NewPin(Name string, Number int, AtRest bool) *Pin {
	// Initialize the board, if not done already
	p := &Pin {
		Name: Name,
		Number: Number,
		AtRest: AtRest,
		Pin: rpio.Pin(Number),
	}
	// Setup the rpio.Pin to be an input pin and set current val
	p.Pin.Input()
	p.Current = rpio.ReadPin(p.Pin) == 1
	return p
}

// Initialize the pins
func Init() {
	// If already initialized, skip this

	if !gpioInit {
		err := rpio.Open()
		if err != nil {
			log.Fatal("Error opening pins", err)
		}

		// Initialize Garage and Front Door Pin
		GaragePin := NewPin("garage", 18, false)
		FrontDoorPin := NewPin("front door", 23, false)
		gpioInit = true
		Pins = append(Pins, GaragePin, FrontDoorPin)
	}
}

// Check on the Pins and pass a status back
func CheckPins(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching pin status")
	Init()

	for i := range Pins {
		p := Pins[i]
		p.Current = rpio.ReadPin(p.Pin) == 1
	}

	// Send back API response
	SetResponseHeaders(w)
	json.NewEncoder(w).Encode(Pins)
}
