package repo

import (
	"time"
)

// User schema
type User struct {
	ID          int32     `gorm:"AUTO_INCREMENT"`
	Stdno       string    `gorm:"type:varchar(10);unique_index;not null"`
	Password    string    `gorm:"type:varchar(64);not null"`
	Realname    string    `gorm:"type:varchar(30);not null"`
	Grade       string    `gorm:"type:varchar(4);not null"`
	College     string    `gorm:"type:varchar(30);not null"`
	Phone       string    `gorm:"type:varchar(11);not null"`
	Qq          string    `gorm:"type:varchar(12);not null"`
	AccessToken string    `gorm:"type:varchar(64);not null"`
	CreateAt    time.Time `gorm:"type:timestamp;not null;"`
	UpdateAt    time.Time `gorm:"type:timestamp;not null;"`
}
