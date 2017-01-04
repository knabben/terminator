package cmd

import (
	restful "github.com/emicklei/go-restful"
	"net/http"
)

func CreateHTTPAPIHandler(http.Handler, error) {

	apiV1 := new(restful.WebService)
	apiV1.Path("/api/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

}
