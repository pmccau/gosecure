package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TPin struct {
	Name 			string
	Number 			int
	AtRest			bool
	Current			bool
	Pin 			interface{}
}

var TPins []*TPin

func TNewPin(Name string, Number int, AtRest bool) *TPin {
	p := &TPin {
		Name: Name,
		Number: Number,
		AtRest: AtRest,
		Pin: nil,
		Current: AtRest,
	}
	return p
}

// Initialize the pins
func TInit() {
	TGaragePin := TNewPin("garage", 18, false)
	TFrontDoorPin := TNewPin("front door", 23, false)
	var tempTPins []*TPin
	TPins = append(tempTPins, TGaragePin, TFrontDoorPin)
}

// Check on the Pins and pass a status back
func TCheckPins(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching pin status TEST")
	TInit()

	for i := range TPins {
		p := TPins[i]
		p.Current = i % 2 == 0
	}

	// Send back API response
	SetResponseHeaders(w)
	json.NewEncoder(w).Encode(TPins)
}