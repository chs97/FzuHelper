package api

import (
	authProto "github.com/chs97/FzuHelper/srv/auth/proto"
	jwchProto "github.com/chs97/FzuHelper/srv/jwch/proto"
	umbProto "github.com/chs97/FzuHelper/srv/umbrella/proto"
	"github.com/micro/go-micro/client"
)

var (
	authClient authProto.AuthClient
	jwchClient jwchProto.JwchClient
	umbClient  umbProto.LendClient
)

func init() {
	authClient = authProto.NewAuthClient("FzuHelper.auth", client.DefaultClient)
	jwchClient = jwchProto.NewJwchClient("FzuHelper.jwch", client.DefaultClient)
	umbClient = umbProto.NewLendClient("FzuHelper.lend", client.DefaultClient)
}
