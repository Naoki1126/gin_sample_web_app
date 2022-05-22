package handler

import (
	"gin_sample_web_app/server/gen/models"
	"gin_sample_web_app/server/gen/restapi/gin_sample_web_app"

	"github.com/go-openapi/runtime/middleware"
)

var (
	Users map[uint64]models.User
)

func init() {
	Users = map[uint64]models.User{}
}

type PostUserHandler struct{}

func (h *PostUserHandler) Handle(params gin_sample_web_app.PostUserParams) middleware.Responder {
	u := models.User{
		ID:   params.Body.ID,
		Name: params.Body.Name,
	}

	Users[params.Body.ID] = u

	return gin_sample_web_app.NewPostUserCreated()
}
