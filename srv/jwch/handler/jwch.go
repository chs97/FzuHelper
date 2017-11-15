package handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/chs97/FzuHelper/srv/jwch/proto"
	"golang.org/x/net/context"
)

type Jwch struct{}

func getClient(stdno, passwd string) (http.Client, string, error) {
	var client http.Client
	jar, err := cookiejar.New(nil)
	if err != nil {
		return client, "", err
	}
	client.Jar = jar
	form := url.Values{"muser": {stdno}, "passwd": {passwd}}
	req, err := http.NewRequest("POST", "http://59.77.226.32/logincheck.asp", strings.NewReader(form.Encode()))
	if err != nil {
		return client, "", err
	}
	req.PostForm = form
	req.Header.Add("Referer", "http://jwch.fzu.edu.cn/")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return client, "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return client, "", err
	}
	result := string(body)
	alert := strings.Index(result, "alert")
	if alert == -1 {
		return client, "", errors.New("Username or Password error")
	}
	start := strings.Index(result, "top.aspx?id=")
	end := strings.Index(result, "\" noresize=\"noresize\"")
	if start == -1 || end == -1 {
		return client, "", errors.New("Id not found")
	}
	id := result[start+len("top.aspx?id=") : end]
	return client, id, nil
}

func (j *Jwch) Login(ctx context.Context, req *jwch.LoginRequest, rsp *jwch.LoginResponse) error {
	stdno := req.Stdno
	password := req.Password
	_, _, err := getClient(stdno, password)
	if err != nil {
		return err
	}
	return nil
}
