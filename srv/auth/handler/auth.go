package handler

import (
	"os"

	"github.com/chs97/FzuHelper/srv/auth/proto"
	"github.com/chs97/FzuHelper/srv/auth/repo"
	"golang.org/x/net/context"
)

const expire = 60

var secret []byte

func init() {
	secret = []byte(os.Getenv("JWT_SECRET"))
}

// Auth Handler
type Auth struct{}

// Create
func (a *Auth) Create(ctx context.Context, req *auth.CreateRequest, rsp *auth.CreateResponse) error {
	user := &repo.User{
		Stdno:    req.Stdno,
		Realname: req.Realname,
		Grade:    req.Grade,
		College:  req.College,
	}
	err := repo.Creat(user, req.Password)
	if err != nil {
		return err
	}
	rsp.Id = user.ID
	return nil
}

func (a *Auth) Read(ctx context.Context, req *auth.ReadRequest, rsp *auth.ReadResponse) error {
	user := new(repo.User)
	err := repo.GetByStdno(req.Stdno, user)
	if err != nil {
		return err
	}
	ret := &auth.User{
		Id:       user.ID,
		Stdno:    user.Stdno,
		Realname: user.Realname,
		College:  user.College,
		Grade:    user.Grade,
	}
	rsp.User = ret
	return nil
}
