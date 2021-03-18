package router

import (
	"github.com/gorilla/mux"
	"gosecure/middleware"
)

// Router's gonna route
func Router() *mux.Router {
	router := mux.NewRouter()

	// Serve
	router.HandleFunc("/api/weather/19147", middleware.GetWeather)
	router.HandleFunc("/api/pins", middleware.CheckPins)
	router.HandleFunc("/api/logs", middleware.CheckPinLogs)
	return router
}
