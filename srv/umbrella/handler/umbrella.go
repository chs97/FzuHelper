package handler

import (
	umb "github.com/chs97/FzuHelper/srv/umbrella/proto"
	"github.com/chs97/FzuHelper/srv/umbrella/repo"
	"golang.org/x/net/context"
)

// Umbrella Handler
type Umbrella struct{}

// GetAllUmb Get all umbrellas
func (u *Umbrella) GetAllUmb(ctx context.Context, req *umb.GetAllUmbRequest, rsp *umb.GetAllumbResponse) error {
	umbrellas := []repo.Umbrella{}
	err := repo.GetAllUmb(&umbrellas)
	if err != nil {
		return err
	}
	records := []*umb.Umbrella{}
	for _, item := range umbrellas {
		record := new(umb.Umbrella)
		record.Number = item.Number
		record.Rented = repo.HasRented(item.Number)
		records = append(records, record)
	}
	rsp.Umbrellas = records
	return nil
}

func (u *Umbrella) BorrowOne(ctx context.Context, req *umb.BorrowOneRequest, rsp *umb.BorrowOneResponse) error {
	return repo.AddBorrowRecord(req.Stdno, req.Place, req.Number)
}

func (u *Umbrella) ReturnOne(ctx context.Context, req *umb.ReturnRequest, rsp *umb.ReturnResponse) error {
	return repo.ReturnUmb(req.Id)
}
func (u *Umbrella) GetRecordsByNum(ctx context.Context, req *umb.GetRecordsByNumRequest, rsp *umb.GetRecordsByNumResponse) error {
	borrows := []repo.Borrow{}
	err := repo.GetRecordsByNumber(req.Number, &borrows)
	if err != nil {
		return err
	}
	records := []*umb.Record{}
	for _, item := range borrows {
		record := new(umb.Record)
		record.Id = item.ID
		record.Stdno = item.Stdno
		record.CreateAt = item.CreateAt.Unix()
		record.HasReturn = item.HasReturn
		record.Place = item.Place
		record.ReturnAt = item.ReturnAt.Unix()
		records = append(records, record)
	}
	rsp.Records = records
	return nil
}

func (u *Umbrella) GetRecordsByStdno(ctx context.Context, req *umb.GetRecordsByStdnoRequest, rsp *umb.GetRecordsByStdnoResponse) error {
	borrows := []repo.Borrow{}
	err := repo.GetRecordsByStdno(req.Stdno, &borrows)
	if err != nil {
		return err
	}
	records := []*umb.Record{}
	for _, item := range borrows {
		record := new(umb.Record)
		record.Id = item.ID
		record.Stdno = item.Stdno
		record.CreateAt = item.CreateAt.Unix()
		record.HasReturn = item.HasReturn
		record.Place = item.Place
		record.ReturnAt = item.ReturnAt.Unix()
		records = append(records, record)
	}
	rsp.Records = records
	return nil
}

func (u *Umbrella) AddUmbrella(ctx context.Context, req *umb.AddUmbrellaRequest, rsp *umb.AddUmbrellaResponse) error {
	return repo.AddUmbrella(req.Number)
}
