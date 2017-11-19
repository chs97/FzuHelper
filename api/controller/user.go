package api

import (
	"errors"

	authProto "github.com/chs97/FzuHelper/srv/auth/proto"
	jwchProto "github.com/chs97/FzuHelper/srv/jwch/proto"
	"github.com/emicklei/go-restful"
	"golang.org/x/net/context"
)

type User struct{}

func (u *User) Login(req *restful.Request, rsp *restful.Response) {
	type reqType struct {
		Stdno  string `json:"stdno"`
		Passwd string `json:"password"`
	}
	var (
		jwchRsp *jwchProto.GetInfoResponse
		jwtRsp  *authProto.JWTSignResponse
		err     error
	)
	res := make(map[string]interface{})
	res["state"] = "0"
	reqData := new(reqType)
	err = req.ReadEntity(reqData)
	if err != nil {
		rsp.WriteError(500, err)
		return
	}
	if len(reqData.Stdno) != 9 {
		rsp.WriteError(400, errors.New("学号的长度必须为9位"))
		return
	}
	if len(reqData.Passwd) == 0 {
		rsp.WriteError(400, errors.New("密码不得为空"))
		return
	}
	_, err = authClient.Read(context.TODO(), &authProto.ReadRequest{
		Stdno: reqData.Stdno,
	})
	if err != nil && err.Error() == "Record not found" {
		jwchRsp, err = jwchClient.Getinfo(context.TODO(), &jwchProto.GetInfoRequest{
			Stdno:    reqData.Stdno,
			Password: reqData.Passwd,
		})
		if err != nil {
			rsp.WriteError(403, err)
			return
		}
		createReq := new(authProto.CreateRequest)
		createReq.College = jwchRsp.College
		createReq.Grade = jwchRsp.Grade
		createReq.Realname = jwchRsp.Realname
		createReq.Password = reqData.Passwd
		createReq.Stdno = jwchRsp.Stdno
		_, err := authClient.Create(context.TODO(), createReq)
		if err != nil {
			rsp.WriteError(500, err)
			return
		}
	} else if err != nil {
		rsp.WriteError(500, err)
		return
	}
	jwtRsp, err = authClient.JWTSign(context.TODO(), &authProto.JWTSignRequest{
		Stdno:    reqData.Stdno,
		Password: reqData.Passwd,
	})
	if err != nil {
		rsp.WriteError(403, err)
		return
	}
	// res["id"] = authRsp.User.Id
	res["state"] = "1"
	res["data"] = jwtRsp.Payload
	rsp.WriteEntity(res)
}

func (u *User) GetInfo(req *restful.Request, rsp *restful.Response) {
	payload := req.HeaderParameter("token")
	JWT, err := authClient.JWTVerify(context.TODO(), &authProto.JWTVerifyRequest{
		Payload: payload,
	})
	if err != nil {
		rsp.WriteError(400, err)
		return
	}
	User, err := authClient.Read(context.TODO(), &authProto.ReadRequest{
		Stdno: JWT.Stdno,
	})
	if err != nil {
		rsp.WriteError(400, err)
		return
	}
	res := make(map[string]interface{})
	data := make(map[string]interface{})
	data["user"] = User.User
	res["data"] = data
	rsp.WriteEntity(res)
}

func (u *User) Update(req *restful.Request, rsp *restful.Response) {
	payload := req.HeaderParameter("token")
	JWT, err := authClient.JWTVerify(context.TODO(), &authProto.JWTVerifyRequest{
		Payload: payload,
	})
	if err != nil {
		rsp.WriteError(400, err)
		return
	}
	stdno := JWT.Stdno
	type reqType struct {
		Phone string `json:"phone"`
		Qq    string `json:"qq"`
	}
	reqData := new(reqType)
	err = req.ReadEntity(reqData)
	if err != nil {
		rsp.WriteError(400, err)
		return
	}
	_, err = authClient.Update(context.TODO(), &authProto.UpdateRequest{
		Stdno: stdno,
		Phone: reqData.Phone,
		Qq:    reqData.Qq,
	})
	if err != nil {
		rsp.WriteError(400, err)
		return
	}
}

func AuthFilter(req *restful.Request, rsp *restful.Response, chain *restful.FilterChain) {
	payload := req.HeaderParameter("token")
	if len(payload) == 0 {
		rsp.WriteError(401, errors.New("Token invalid"))
		return
	}
	_, err := authClient.JWTVerify(context.TODO(), &authProto.JWTVerifyRequest{
		Payload: payload,
	})
	if err != nil {
		rsp.WriteError(401, err)
		return
	}
	chain.ProcessFilter(req, rsp)
}

var UserService *restful.WebService

func init() {
	user := new(User)
	UserService = new(restful.WebService)
	UserService.Consumes(restful.MIME_JSON, restful.MIME_JSON)
	UserService.Produces(restful.MIME_JSON, restful.MIME_JSON)
	UserService.Path("/user")
	UserService.Route(UserService.POST("/login").To(user.Login))
	UserService.Route(UserService.GET("/").Filter(AuthFilter).To(user.GetInfo))
	UserService.Route(UserService.PUT("/").Filter(AuthFilter).To(user.Update))
}
