package api

import (
	umbProto "github.com/chs97/FzuHelper/srv/umbrella/proto"
	restful "github.com/emicklei/go-restful"
	"golang.org/x/net/context"
)

// Lend API
type Lend struct{}

func (l *Lend) BorrowOne(req *restful.Request, rsp *restful.Response) {
	type reqType struct {
		Place  string
		Number string
	}
	reqData := new(reqType)
	err := req.ReadEntity(reqData)
	res := make(map[string]interface{})
	res["state"] = "0"
	defer rsp.WriteEntity(res)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	_, err = umbClient.BorrowOne(context.TODO(), &umbProto.BorrowOneRequest{
		Stdno:  TStdno,
		Place:  reqData.Place,
		Number: reqData.Number,
	})
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["state"] = "1"
}

func (l *Lend) ReturnOne(req *restful.Request, rsp *restful.Response) {

}

func (l *Lend) GetAllMyBorrow(req *restful.Request, rsp *restful.Response) {
	res := make(map[string]interface{})
	defer rsp.WriteEntity(res)
	res["state"] = "0"
	records, err := umbClient.GetRecordsByStdno(context.TODO(), &umbProto.GetRecordsByStdnoRequest{
		Stdno: TStdno,
	})
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["data"] = records.Records
	res["state"] = "1"
}

var UmbService *restful.WebService

func init() {
	lend := new(Lend)
	UmbService = new(restful.WebService)
	UmbService.Consumes(restful.MIME_JSON, restful.MIME_JSON)
	UmbService.Produces(restful.MIME_JSON, restful.MIME_JSON)
	UmbService.Path("/api/umbrella")
	UmbService.Route(UmbService.POST("/borrow").Filter(AuthFilter).To(lend.BorrowOne))
	UmbService.Route(UmbService.GET("/").Filter(AuthFilter).To(lend.GetAllMyBorrow))
}
