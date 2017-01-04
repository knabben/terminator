package cmd

import (
	restful "github.com/emicklei/go-restful"
)

type User struct {
	Id, Name string
}

func CreateHTTPAPIHandler() *restful.WebService {
	apiV1 := new(restful.WebService)
	apiV1.Path("/api/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	apiV1.Route(apiV1.GET("/test").To(TestF))

	return apiV1
}

func TestF(request *restful.Request, response *restful.Response) {
	response.WriteEntity("test")
}
