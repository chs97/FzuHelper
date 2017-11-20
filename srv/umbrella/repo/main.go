package repo

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func init() {
	host := os.Getenv("POSTGRES_HOST")
	database := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")

	config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, database, password)

	if host == "" || database == "" {
		log.Fatal("should provide POSTGRES_HOST and POSTGRES_DB")
	}

	var err error

	db, err = gorm.Open("postgres", config)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Borrow{})
	db.AutoMigrate(&Umbrella{})
}

func AddUmbrella(number string) error {
	umb := new(Umbrella)
	res := db.Where("number = ?", number).First(umb)
	if !res.RecordNotFound() {
		return errors.New("Number has exist")
	}
	umb.Number = number
	return db.Create(umb).Error
}

func AddBorrowRecord(stdno, place, number string) error {
	borrow := new(Borrow)
	ret := db.Where("stdno = ? AND has_return = ?", stdno, true).First(&Borrow{})
	if !ret.RecordNotFound() {
		return errors.New("Borrow Not Return")
	}
	retM := db.Where("number = ? AND has_return = ?", number, true).First(&Borrow{})
	if !retM.RecordNotFound() {
		return errors.New("Umbrella Has Rented")
	}
	borrow.Stdno = stdno
	borrow.HasReturn = false
	borrow.CreateAt = time.Now()
	borrow.Place = place
	borrow.Number = number
	return db.Create(borrow).Error
}

func ReturnUmb(id int32) error {
	borrow := new(Borrow)
	ret := db.Where("id = ?", id).First(borrow)
	if ret.RecordNotFound() {
		return errors.New("Record not found")
	}
	borrow.ReturnAt = time.Now()
	borrow.HasReturn = true
	return db.Save(borrow).Error
}

func HasRented(number string) bool {
	borrow := new(Borrow)
	ret := db.Where("number = ? AND has_return = ?", number, false).First(borrow)
	if ret.RecordNotFound() {
		return false
	}
	return true
}

func GetRecordsByNumber(number string, records []Borrow) error {
	ret := db.Where("number = ? AND has_return = ?", number, false).Find(records)
	if ret.RecordNotFound() {
		return errors.New("Record not found")
	}
	return nil
}

func GetRecordsByStdno(stdno string, records []Borrow) error {
	ret := db.Where("stdno = ? ", stdno).Find(records)
	if ret.RecordNotFound() {
		return errors.New("Record not found")
	}
	return nil
}

func GetAllUmb(umbrellas []Umbrella) error {
	ret := db.Find(umbrellas)
	return ret.Error
}
