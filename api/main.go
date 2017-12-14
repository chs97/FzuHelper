package main

import (
	"os"

	"github.com/chs97/FzuHelper/api/controller"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.PartyFunc("/user", api.UserController)
	app.WrapRouter(cors.WrapNext(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PATCH", "PUT"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	app.Run(iris.Addr(":"+port), iris.WithoutVersionChecker)
}
