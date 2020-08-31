package router

import (
	"github.com/gorilla/mux"
	"gosecure/middleware"
)

// Router's gonna route
func Router() *mux.Router {
	router := mux.NewRouter()

	// Serve
	router.HandleFunc("/api/weather", middleware.GetWeather)
	router.HandleFunc("/api/pins", middleware.CheckPins)
	return router
}
