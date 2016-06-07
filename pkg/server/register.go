package server

import (
	"../api/weather"
	"github.com/emicklei/go-restful"

	"log"
	"net/http"
)

func Serve(port string) {

	wsContainer := restful.NewContainer()

	wr := WeatherResource{
		WeatherData: []weather.WeatherPoint{
			weather.WeatherPoint{
				ID: 10,
				Location: weather.Location{
					City:       "Jakobstad",
					PostalCode: 68600,
				},
				Temperature: 20,
				Humidity:    55,
			},
		},
	}

	wr.Register(wsContainer)

	registerVersion(wsContainer)

	log.Printf("start listening on localhost:" + port)

	server := &http.Server{Addr: ":" + port, Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
