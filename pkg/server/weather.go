package server

import (
	"../api/weather"
	"github.com/emicklei/go-restful"

	"net/http"
	"strconv"
)

type WeatherResource struct {
	WeatherData []weather.WeatherPoint
}

func (w WeatherResource) Register(container *restful.Container) {
	ws := new(restful.WebService)

	ws.Path("/weather").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/").To(w.findAllWeatherData).
		Operation("findAllWeatherData").
		Writes([]weather.WeatherPoint{}))

	ws.Route(ws.GET("/{weather-id}").To(w.findWeatherData).
		Operation("findWeatherData").
		Param(ws.PathParameter("weather-id", "identifier of the weather point").DataType("int32")).
		Writes(weather.WeatherPoint{}))

	container.Add(ws)
}

func (w WeatherResource) findAllWeatherData(request *restful.Request, response *restful.Response) {
	response.WriteEntity(w.WeatherData)
}

func (w WeatherResource) findWeatherData(request *restful.Request, response *restful.Response) {
	str_id := request.PathParameter("weather-id")
	id,_ := strconv.Atoi(str_id)
	var weatherPoint *weather.WeatherPoint

	for _,p := range w.WeatherData {
		if p.ID == int32(id) {
			weatherPoint = &p
		}
	}

	if weatherPoint == nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "404: That weather point could not be found.")
		return
	}

	response.WriteEntity(weatherPoint)
}
