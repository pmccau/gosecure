package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// GetWeather will pass a request along to the open weather API, then return a forecast
func GetWeather(w http.ResponseWriter, r *http.Request) {
	var res interface{}
	fmt.Println("Fetching weather")
	api_key := readCredFromFile("./cred.pickle")

	// Retrieve the data from the weather site
	_ = json.NewDecoder(r.Body).Decode(&res)
	var client = &http.Client{Timeout: 10*time.Second}
	apiRes, err := client.Get("https://api.openweathermap.org/data/2.5/weather?zip=19147&appid=" + api_key)
	if err != nil {
        fmt.Println(err)
	}
	defer apiRes.Body.Close()

	// Parse API response
	var temp interface{}
	json.NewDecoder(apiRes.Body).Decode(&temp)
	jsonStr, err := json.Marshal(temp)
	var parsedResponse map[string]interface{}
	err = json.Unmarshal(jsonStr, &parsedResponse)

	// Respond to the requester (pass through)
	SetResponseHeaders(w)
	json.NewEncoder(w).Encode(parsedResponse)
}

// prettyPrint is a debugging method for printing json to console
func prettyPrint(a map[string] interface{}) {
	b, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(string(b))
}

// SetResponseHeaders will set the headers appropriately before sending back a response
// for all middleware components
func SetResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Context-Type", "application/results-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func readCredFromFile(filename string) string {
	// Credentials
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(dat))
}
