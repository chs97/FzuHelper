package api

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"

	authProto "github.com/chs97/FzuHelper/srv/auth/proto"
	jwchProto "github.com/chs97/FzuHelper/srv/jwch/proto"
	"golang.org/x/net/context"
)

// User API
type User struct{}

// TStdno auth -> stdno
// var TStdno = ""

func userLogin(ctx iris.Context) {
	type reqType struct {
		Stdno    string
		Password string
	}
	res := make(map[string]interface{})
	reqData := new(reqType)
	ctx.ReadJSON(reqData)
	if len(reqData.Stdno) != 9 {
		ctx.StatusCode(httptest.StatusBadRequest)
		ctx.Text("学号的长度必须为9位")
		return
	}
	if len(reqData.Password) == 0 {
		ctx.StatusCode(httptest.StatusBadRequest)
		ctx.Text("密码不得为空")
		return
	}
	jwtSignRes, err := authClient.JWTSign(context.TODO(), &authProto.JWTSignRequest{
		Stdno:    reqData.Stdno,
		Password: reqData.Password,
	})
	// login failure password error
	if err != nil && err.Error() == "Password error" {
		_, err := jwchClient.Getinfo(context.TODO(), &jwchProto.GetInfoRequest{
			Stdno:    reqData.Stdno,
			Password: reqData.Password,
		})
		if err != nil {
			fmt.Println(err.Error())
			ctx.StatusCode(httptest.StatusBadRequest)
			ctx.Text("账号或密码错误")
			return
		}
		changePwd, err := authClient.ChangePwd(context.TODO(), &authProto.ChangePwdRequest{
			Stdno:    reqData.Stdno,
			Password: reqData.Password,
		})
		if err != nil {
			ctx.StatusCode(httptest.StatusInternalServerError)
			ctx.Text(err.Error())
			return
		}
		res["data"] = changePwd.Payload
		ctx.JSON(res)
		return
	}
	if err == nil {
		res["data"] = jwtSignRes.Payload
		ctx.JSON(res)
		return
	}

	if err != nil && err.Error() != "Record not found" {
		ctx.StatusCode(httptest.StatusInternalServerError)
		return
	}
	// record not found then create user
	jwchReq, err := jwchClient.Getinfo(context.TODO(), &jwchProto.GetInfoRequest{
		Stdno:    reqData.Stdno,
		Password: reqData.Password,
	})
	if err != nil {
		ctx.StatusCode(httptest.StatusInternalServerError)
		ctx.Text(err.Error())
		return
	}
	createRes, err := authClient.Create(context.TODO(), &authProto.CreateRequest{
		Stdno:    reqData.Stdno,
		Password: reqData.Password,
		Grade:    jwchReq.Grade,
		College:  jwchReq.College,
		Realname: jwchReq.Realname,
	})
	if err != nil {
		ctx.StatusCode(httptest.StatusInternalServerError)
		ctx.Text(err.Error())
		return
	}
	res["data"] = createRes.Token
	ctx.JSON(res)
	return
}

func getUserInfo(ctx iris.Context) {
	stdno := ctx.Values().GetString("stdno")
	User, err := authClient.Read(context.TODO(), &authProto.ReadRequest{
		Stdno: stdno,
	})
	if err != nil {
		ctx.StatusCode(400)
		ctx.Text(err.Error())
		return
	}
	res := make(map[string]interface{})
	data := make(map[string]interface{})
	data["user"] = User.User
	res["data"] = data
	ctx.JSON(res)
}

// UserController user api
func UserController(user iris.Party) {
	user.Get("/", Authentication, getUserInfo)
	user.Post("/", userLogin)
}
