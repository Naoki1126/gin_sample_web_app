package handler

import (
	"strconv"

	"github.com/go-openapi/runtime/middleware"

	"gin_sample_web_app/server/gen/restapi/gin_sample_web_app"
)

type GetUserByIDHandler struct{}

func (h *GetUserByIDHandler) Handle(params gin_sample_web_app.GetUserByIDParams) middleware.Responder {
	id, _ := strconv.ParseUint(params.ID, 10, 64)

	u := Users[id]

	return gin_sample_web_app.NewGetUserByIDOK().WithPayload(&u)
}
