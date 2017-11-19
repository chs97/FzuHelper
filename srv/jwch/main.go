package main

import (
	"log"

	"github.com/chs97/FzuHelper/srv/jwch/handler"
	jwchProto "github.com/chs97/FzuHelper/srv/jwch/proto"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("FzuHelper.jwch"),
		micro.Version("latest"),
	)

	service.Init()

	jwchProto.RegisterJwchHandler(service.Server(), new(handler.Jwch))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
