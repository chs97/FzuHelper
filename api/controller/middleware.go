package api

import (
	authProto "github.com/chs97/FzuHelper/srv/auth/proto"
	"github.com/kataras/iris"
	"golang.org/x/net/context"
)

// Authentication middleware set stdno to ctx
func Authentication(ctx iris.Context) {
	payload := ctx.GetHeader("Authorization")
	if len(payload) == 0 {
		ctx.StatusCode(401)
		ctx.Text("Token invalid")
		return
	}
	token, err := authClient.JWTVerify(context.TODO(), &authProto.JWTVerifyRequest{
		Payload: payload,
	})
	if err != nil {
		ctx.StatusCode(401)
		ctx.Text(err.Error())
		return
	}
	ctx.Values().Set("stdno", token.Stdno)
	ctx.Next()
}
