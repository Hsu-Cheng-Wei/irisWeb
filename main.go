package main

import (
	"irisWeb/controller/controllerinit"

	_ "irisWeb/controller"
)

func main() {
	app := controllerinit.AppRouter.App

	app.Listen(":8080")
}
