package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type TPin struct {
	Name 			string
	Number 			int
	AtRest			bool
	Current			bool
	Pin 			interface{}
}

var TLOG_FILE = "logs/pins_test.log"
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
		TLogPinEvent(p)
	}

	// Send back API response
	SetResponseHeaders(w)
	json.NewEncoder(w).Encode(TPins)
}

// LogPinEvent will do some simple logging about a status change in a Pin
func TLogPinEvent(p *TPin) {
	timestamp := time.Now().Format("Mon Jan _2 15:04:05 2006")
	toWrite := fmt.Sprintf("[%s] Detected shift in %s from %s to %s\n",
		timestamp, p.Name, PinStatusToString(p.Current), PinStatusToString(!p.Current))

	f, err := os.OpenFile(TLOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()
	if _, err = f.WriteString(toWrite); err != nil {
		log.Fatal(err)
	}
}

// Should read the log's tail, then serve it up
func TReadLog() string {
	f, err := os.OpenFile(TLOG_FILE, os.O_RDONLY, 0666)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	var bytesTail int64 = 1000
	buf := make([]byte, bytesTail)
	stat, err := os.Stat(TLOG_FILE)
	start := stat.Size() - bytesTail
	_, err = f.ReadAt(buf, start)
	if err != nil {
		log.Fatal("Error serving up log")
	}
	return string(buf)
}

// Check on the Pins and pass a status back
func TCheckPinLogs(w http.ResponseWriter, r *http.Request) {
	res := make(map[string]string)
	res["events"] = TReadLog()

	// Send back API response
	SetResponseHeaders(w)
	json.NewEncoder(w).Encode(res)
}