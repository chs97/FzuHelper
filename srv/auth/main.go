package main

import (
	"log"

	"github.com/chs97/FzuHelper/srv/auth/handler"
	authProto "github.com/chs97/FzuHelper/srv/auth/proto"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("FzuHelper.auth"),
		micro.Version("latest"),
	)

	service.Init()

	authProto.RegisterAuthHandler(service.Server(), new(handler.Auth))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
