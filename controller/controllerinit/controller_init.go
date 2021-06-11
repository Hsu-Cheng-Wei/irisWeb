package controllerinit

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

type Router struct {
	App *iris.Application

	V1Party router.Party

	JwtMiddleware *jwt.Middleware

	JwtToken string
}

const jwt_token = "00447fb4-9a87-430d-bece-13565d740b7d"

var AppRouter Router

func (r Router) GetJwtSignedToken(key jwt.MapClaims) string {
	claim := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, key)

	token, _ := claim.SignedString([]byte(jwt_token))

	return token
}

func init() {
	newRouter()
}

func newRouter() {
	AppRouter.App = iris.New()

	AppRouter.V1Party = AppRouter.App.Party("/api/v1")

	AppRouter.JwtToken = jwt_token

	AppRouter.JwtMiddleware = jwtConfig()
}

func jwtConfig() *jwt.Middleware {
	return jwt.New(jwt.Config{

		Extractor: jwt.FromAuthHeader,

		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(jwt_token), nil
		},

		SigningMethod: jwt.SigningMethodHS256,
	})
}
