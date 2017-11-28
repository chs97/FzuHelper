package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chs97/FzuHelper/srv/jwch/proto"
	"golang.org/x/net/context"
)

type Jwch struct{}

type ocr struct {
	Base64    string `json:"base64"`
	Trim      string `json:"trim"`
	whitelist string `json:"whitelist"`
}

type ocrResult struct {
	Result  string `json:"result"`
	Version string `json::"version"`
}

func captcha(base64 string) string {
	data := ocr{Base64: base64, Trim: "\n", whitelist: "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(data)
	res, err := http.Post("http://soft.hs97.cn:8088/base64", "application/json; charset=utf-8", buf)
	if err != nil {
		return ""
	}
	resData := new(ocrResult)
	json.NewDecoder(res.Body).Decode(resData)
	return resData.Result
}

func getClient(stdno, passwd string) (http.Client, string, error) {
	var client http.Client
	jar, err := cookiejar.New(nil)
	if err != nil {
		return client, "", err
	}
	client.Jar = jar
	getCapReq, err := http.NewRequest("GET", "http://59.77.226.32/captcha.asp?time=0.9019912959584149", nil)
	if err != nil {
		return client, "", err
	}
	getCapRes, err := client.Do(getCapReq)
	if err != nil {
		return client, "", err
	}
	capBody, err := ioutil.ReadAll(getCapRes.Body)
	if err != nil {
		return client, "", err
	}
	base := base64.StdEncoding.EncodeToString(capBody)
	capResult := captcha(base)
	form := url.Values{"muser": {stdno}, "passwd": {passwd}, "Verifycode": {capResult}}
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
	if len(result) == 193 {
		return client, "", errors.New("Username or Password error")
	}
	if len(result) == 167 {
		return client, "", errors.New("Captcha Error")
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

// Getinfo asdasdsa
func (j *Jwch) Getinfo(ctx context.Context, req *jwch.GetInfoRequest, rsp *jwch.GetInfoResponse) error {
	stdno := req.Stdno
	passwd := req.Password
	var client http.Client
	var id string
	var err error
	for {
		client, id, err = getClient(stdno, passwd)
		if err != nil && err.Error() == "Captcha Error" {
			continue
		}
		break
	}
	if err != nil {
		return err
	}

	url := "http://59.77.226.35/jcxx/xsxx/StudentInformation.aspx?id=" + id
	jreq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	jreq.Header.Add("Referer", "http://jwch.fzu.edu.cn/")
	resp, err := client.Do(jreq)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	rsp.Stdno = doc.Find("#ContentPlaceHolder1_LB_xh").Text()
	rsp.Realname = doc.Find("#ContentPlaceHolder1_LB_xm").Text()
	rsp.College = doc.Find("#ContentPlaceHolder1_LB_xymc").Text()
	rsp.Grade = doc.Find("#ContentPlaceHolder1_LB_nj").Text()
	return nil
}
