package user

import (
	"irisWeb/controller/controllerinit"
	"irisWeb/domain"
	"irisWeb/service"

	"github.com/google/uuid"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

func GetAll(userService service.UserService) []UserDto {
	users := userService.GetAll()
	result := []UserDto{}
	for _, user := range users {

		result = append(result,&UserDto{
			ID: uuid.
		})
	}
}

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

func InserUser(userService service.UserService, user domain.User) string {
	id, err := userService.Insert(user)

	if err != nil {
		panic(err)
	}

	uid, _ := uuid.ParseBytes(id)

	return uid.String()
}

func DeleteUser(userService service.UserService, id string) error {
	return userService.Delete(id)
}
