package server

import (
	"gin_sample_web_app/server/gen/restapi/gin_sample_web_app"

	"github.com/go-openapi/runtime/middleware"
)

func GetGreeting(p gin_sample_web_app.GetGreetingParams) middleware.Responder {
	payload := *p.Name
	return gin_sample_web_app.NewGetGreetingOK().WithPayload(payload)
}
