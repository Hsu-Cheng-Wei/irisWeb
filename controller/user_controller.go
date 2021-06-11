package controller

import (
	"irisWeb/controller/controllerinit"
	"irisWeb/database"
	"irisWeb/handler/user"
	"irisWeb/repository"
	"irisWeb/service"

	"github.com/kataras/iris/v12"
)

func init() {
	api := controllerinit.AppRouter.V1Party.Party("/user")

	api.RegisterDependency(func(ctx iris.Context) service.UserService {
		return &repository.UserRepository{
			Db: database.MysqlOrm,
		}
	})

	api.ConfigureContainer(containerCfg)
}

func containerCfg(api *iris.APIContainer) {
	j := controllerinit.AppRouter.JwtMiddleware

	api.Get("/all", user.GetAll)
	api.Post("/signin", user.Auth)
	api.Post("/signup", j.Serve, user.InserUser)
	api.Delete("/", j.Serve, user.DeleteUser)
}
