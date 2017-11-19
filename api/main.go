package main

import (
	"log"

	"github.com/chs97/FzuHelper/api/controller"
	"github.com/emicklei/go-restful"
	"github.com/micro/go-web"
)

func main() {
	service := web.NewService(
		web.Name("go.micro.api.user"),
	)

	service.Init()

	webContain := restful.NewContainer()
	webContain.Add(api.UserService)
	service.Handle("/", webContain)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
