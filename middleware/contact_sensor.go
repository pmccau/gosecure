package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"log"
	"net/http"
	"os"
	"time"
)

type Pin struct {
	Name 			string
	Number 			int
	AtRest			bool
	Current			bool
	Pin 			rpio.Pin
}

var LOG_FILE = "logs/pins.log"
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
//	fmt.Println("Fetching pin status")
	Init()
	var name string
	var toWrite string
	var currentStateStr string
	var stateChanged []*Pin

	for i := range Pins {
		p := Pins[i]
		nextState := rpio.ReadPin(p.Pin) != 1
		if (nextState != p.Current) {
			stateChanged = append(stateChanged, p)
		}
		p.Current = nextState
	}

	// Send back API response
	SetResponseHeaders(w)
	json.NewEncoder(w).Encode(Pins)

	// Send the email afterwards to not delay the sound
	for i := range stateChanged {
		name, currentStateStr, toWrite = LogPinEvent(stateChanged[i])
		SendMail(fmt.Sprintf("%s is %s", name, currentStateStr), toWrite)
		fmt.Println("Changed!")
	}
}

// Check on the Pins and pass a status back
func CheckPinLogs(w http.ResponseWriter, r *http.Request) {
	res := make(map[string]string)
	res["events"] = ReadLog()

	// Send back API response
	SetResponseHeaders(w)
	json.NewEncoder(w).Encode(res)
}

// LogPinEvent will do some simple logging about a status change in a Pin
func LogPinEvent(p *Pin) (string, string, string) {
	timestamp := time.Now().Format("Mon Jan _2 15:04:05 2006")
	currentStateStr := PinStatusToString(p.Current)
	futureStateStr := PinStatusToString(!p.Current)
	toWrite := fmt.Sprintf("[%s] Detected shift in %s from %s to %s\n",
		timestamp, p.Name, currentStateStr, futureStateStr)

	f, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666);
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err = f.WriteString(toWrite); err != nil {
		log.Fatal(err)
	}
	return p.Name, currentStateStr, toWrite
}

// Should read the log's tail, then serve it up
func ReadLog() string {
	f, err := os.OpenFile(LOG_FILE, os.O_RDONLY, 0666)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	var bytesTail int64 = 1000
	buf := make([]byte, bytesTail)
	stat, err := os.Stat(LOG_FILE)
	if err != nil {
		log.Fatal("Error stat'ing log file", err)
	}
	start := stat.Size() - bytesTail
	_, err = f.ReadAt(buf, start)
	if err != nil {
		log.Fatal("Error serving up log", err)
	}
	return string(buf)
}

func PinStatusToString(status bool) string {
	out := "closed"
	if status {
		out = "open"
	}
	return out
}
