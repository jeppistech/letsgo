package server

import (
	"../version"
	"github.com/emicklei/go-restful"
)

func registerVersion(container *restful.Container) {
	ws := new(restful.WebService)

	ws.Path("/version").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/").To(getVersion).
		Operation("getVersion").
		Writes(""))

	container.Add(ws)
}

func getVersion(request *restful.Request, response *restful.Response) {
	response.WriteEntity(version.Version)
}
