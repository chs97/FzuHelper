package handler

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/chs97/FzuHelper/srv/auth/proto"
	"github.com/chs97/FzuHelper/srv/auth/repo"
	"github.com/chs97/FzuHelper/utils"
	"golang.org/x/net/context"
)

const expire = 12 * 60 * 60

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
		Phone:    user.Phone,
		Qq:       user.Qq,
	}
	rsp.User = ret
	return nil
}

func (a *Auth) Update(ctx context.Context, req *auth.UpdateRequest, rsp *auth.UpdateResponse) error {
	return repo.Update(req.Stdno, req.Phone, req.Qq)
}

func (a *Auth) JWTSign(ctx context.Context, req *auth.JWTSignRequest, rsp *auth.JWTSignResponse) error {
	user := new(repo.User)
	err := repo.GetByStdno(req.Stdno, user)
	if err != nil {
		return err
	}
	if user.Password != utils.Hash(req.Password+repo.Salt) {
		return errors.New("Password error")
	}

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + expire,
		Subject:   user.Stdno,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secret)
	if err != nil {
		return err
	}
	rsp.Payload = ss
	return nil
}

func (a *Auth) JWTVerify(ctx context.Context, req *auth.JWTVerifyRequest, rsp *auth.JWTverifyResponse) error {
	ss := req.Payload
	token, err := jwt.ParseWithClaims(ss, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		return errors.New("Token invalid")
	}
	if err != nil {
		return err
	}
	rsp.Stdno = claims.Subject
	return nil
}
