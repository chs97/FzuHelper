package main

import (
	"log"

	"github.com/chs97/FzuHelper/srv/umbrella/handler"
	umbrellaProto "github.com/chs97/FzuHelper/srv/umbrella/proto"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("FzuHelper.lend"),
		micro.Version("latest"),
	)

	service.Init()

	umbrellaProto.RegisterLendHandler(service.Server(), new(handler.Umbrella))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
