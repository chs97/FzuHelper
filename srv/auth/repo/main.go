package repo

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chs97/FzuHelper/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driven
)

var (
	db *gorm.DB
	// Salt password salt
	Salt string
)

func init() {
	host := os.Getenv("POSTGRES_HOST")
	database := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	Salt = os.Getenv("SALT")

	config := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, database, password)

	if host == "" || database == "" {
		log.Fatal("should provide POSTGRES_HOST and POSTGRES_DB")
	}
	fmt.Println(config)
	var err error

	if Salt == "" {
		log.Fatal("should provide Salt")
	}

	db, err = gorm.Open("postgres", config)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})
}

// Creat User
func Creat(user *User, password string) error {
	// user := new(User)
	timestamp := time.Now()

	user.Password = utils.Hash(password + Salt)
	user.AccessToken = utils.RandomHex(32)
	user.CreateAt = timestamp
	user.UpdateAt = timestamp
	return db.Create(user).Error
}

// GetByStdno get user infomation by stdno
func GetByStdno(stdno string, user *User) error {
	res := db.Where("stdno = ?", stdno).First(user)
	if res.RecordNotFound() {
		return errors.New("Record not found")
	}
	return res.Error
}

// Update update user's infomation
func Update(stdno, phone, qq string) error {
	user := new(User)
	ret := db.Where("stdno = ?", stdno).First(user)
	if ret.RecordNotFound() {
		return errors.New("Record not found")
	}
	if ret.Error != nil {
		return ret.Error
	}
	user.Phone = phone
	user.Qq = qq
	err := db.Save(user)
	return err.Error
}

// ChangePwd change user password
func ChangePwd(stdno, password string) error {
	user := new(User)
	ret := db.Where("stdno = ?", stdno).First(user)
	if ret.RecordNotFound() {
		return errors.New("User not found")
	}
	if ret.Error != nil {
		return ret.Error
	}
	user.Password = utils.Hash(password + Salt)
	return db.Save(user).Error
}
