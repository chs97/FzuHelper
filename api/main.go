package main

import (
	"os"

	"github.com/chs97/FzuHelper/api/controller"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.WrapRouter(cors.WrapNext(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PATCH", "PUT"},
		AllowCredentials: true,
	}))
	app.PartyFunc("/user", api.UserController)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	app.Run(iris.Addr(":" + port))
}
