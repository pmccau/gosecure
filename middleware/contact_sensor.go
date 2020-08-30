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

var TEST_MODE = true

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

func NewTestPin(Name string, Number int, AtRest bool, Current bool) *Pin {
	p := &Pin {
		Name: Name,
		Number: Number,
		AtRest: AtRest,
		Pin: nil,
		Current: AtRest,
	}
	return p
}

// Initialize the pins
func Init() {
	// If already initialized, skip this

	if !TEST_MODE {
		if !gpioInit {
			err := rpio.Open()
			if err != nil {
				log.Fatal("Error opening pins", err)
			}

			// Initialize Garage and Front Door Pin
			GaragePin := NewPin("garage", 18, false)
			FrontDoorPin := NewPin("front door", 23, false)

			Pins = append(Pins, GaragePin, FrontDoorPin)
			for p := range Pins {
				Pins[p].Pin.Detect(rpio.RiseEdge)
			}
			gpioInit = true
		}
	} else {
		GaragePin := NewTestPin("garage", 18, false)
		FrontDoorPin := NewTestPin("front door", 23, false)
		Pins = append(Pins, GaragePin, FrontDoorPin)
	}


}

// Check on the Pins and pass a status back
func CheckPins(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching pin status")
	Init()

	if !TEST_MODE {
		// Retrieve data from the GPIO
		for i := range Pins {
			p := Pins[i]
			p.Current = p.Pin.EdgeDetected()
		}
	} else {
		for i := range Pins {
			p := Pins[i]
			p.Current = i % 2 == 0
		}
	}

	// Send back API response
	SetResponseHeaders(w)
	json.NewEncoder(w).Encode(Pins)
}