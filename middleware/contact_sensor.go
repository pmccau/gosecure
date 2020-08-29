package middleware

//import (
//	"encoding/json"
//	"fmt"
//	"github.com/stianeikeland/go-rpio"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"strings"
//	"time"
//)
//
//type Pin struct {
//	Name 			string
//	Number 			int
//	AtRest			bool
//	Current			bool
//	Pin 			rpio.Pin
//}
//
//var gpioInit bool = false
//var Pins []*Pin
//
//// NewPin is the constructor for a Pin
//func NewPin(Name string, Number int, AtRest bool) *Pin {
//	// Initialize the board, if not done already
//	if !gpioInit {
//		err := rpio.Open()
//		if err != nil {
//			log.Fatal("Error opening pins", err)
//		}
//	}
//
//	p := &Pin {
//		Name: Name,
//		Number: Number,
//		AtRest: AtRest,
//		Pin: rpio.Pin(Number),
//	}
//	// Setup the rpio.Pin to be an input pin and set current val
//	p.Pin.Input()
//	p.Current = rpio.ReadPin(p.Pin) == 1
//	return p
//}
//
//func Init() {
//	GaragePin := NewPin("garage", 18, false)
//	FrontDoorPin := NewPin("front door", 23, false)
//
//	Pins = append(Pins, GaragePin, FrontDoorPin)
//	for p := range Pins {
//		Pins[p].Pin.Detect(rpio.RiseEdge)
//	}
//}
//
//
//func CheckPin(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("Fetching pin status")
//
//	// Retrieve data from the GPIO
//	for p := range Pins {
//		if Pins[p].Pin.EdgeDetected() {
//
//		}
//	}
//
//	// Send back API response
//	SetResponseHeaders(w)
//	json.NewEncoder(w).Encode()
//}

///////////////////////////////////////////////////////////////////////

//func GetWeather(w http.ResponseWriter, r *http.Request) {
//	var res interface{}
//	fmt.Println("Fetching weather")
//
//	// Credentials
//	dat, err := ioutil.ReadFile("./cred.pickle")
//	if err != nil {
//		log.Fatal(err)
//	}
//	api_key := strings.TrimSpace(string(dat))
//
//	// Retrieve the data from the weather site
//	_ = json.NewDecoder(r.Body).Decode(&res)
//	var client = &http.Client{Timeout: 10*time.Second}
//	apiRes, err := client.Get("https://api.openweathermap.org/data/2.5/weather?zip=19147&appid=" + api_key)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer apiRes.Body.Close()
//
//	// Parse API response
//	var temp interface{}
//	json.NewDecoder(apiRes.Body).Decode(&temp)
//	jsonStr, err := json.Marshal(temp)
//	var parsedResponse map[string]interface{}
//	err = json.Unmarshal(jsonStr, &parsedResponse)
//
//	// Respond to the requester (pass through)
//	SetResponseHeaders(w)
//	json.NewEncoder(w).Encode(parsedResponse)
//}