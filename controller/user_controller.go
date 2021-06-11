package controller

import (
	"irisWeb/controller/controllerinit"
	"irisWeb/handler"
)

func init() {
	api := controllerinit.AppRouter.V1Party

	j := controllerinit.AppRouter.JwtMiddleware

	api.Post("/login", handler.Auth)

	api.Post("/user", j.Serve, handler.InserUser)
}
