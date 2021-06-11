package handler

import (
	"irisWeb/controller/controllerinit"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

func Auth(ctx iris.Context) {
	token := controllerinit.AppRouter.GetJwtSignedToken(jwt.MapClaims{
		"user":     "james",
		"password": "123456",
	})

	ctx.JSON(iris.Map{
		"status":  200,
		"message": "success",
		"token":   token,
	})
}

func InserUser(ctx iris.Context) {
	ctx.WriteString("Insert user")
}
