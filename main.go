package main

import (
	"log"
	"net/http"

	"github.com/diwakarydv37/handlers"
	"github.com/diwakarydv37/mux"
	"github.com/diwakarydv37/weather-api/config"
	"github.com/diwakarydv37/weather-api/controller"
	v2 "github.com/diwakarydv37/weather-api/controller/v2"
)

func main() {

	weather := mux.NewRouter()
	weather.Path("/weather/{city}").Methods(http.MethodGet).HandlerFunc(controller.CurrentWeather)

	weather.
		Path("/v2/weather/{city}").
		Queries("unit", "{unit}").
		Methods(http.MethodGet).
		HandlerFunc(v2.CurrentWeather)

	weather.
		Path("/v2/weather/{city}").
		Methods(http.MethodGet).
		HandlerFunc(v2.CurrentWeather)

	if err := http.ListenAndServe(":"+config.Get().Port, handlers.CORS()(weather)); err != nil {
		log.Fatal(err)
	}

}