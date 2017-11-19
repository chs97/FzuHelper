package repo

import "time"

// Borrow schema
type Borrow struct {
	ID        int32     `gorm:"AUTO_INCREMENT"`
	Stdno     string    `gorm:"type:varchar(10)"`
	Place     string    `gorm:"type:varchar(30)"`
	Number    string    `gorm:"type:varchar(10)"`
	CreateAt  time.Time `gorm:"type:timestamp"`
	HasReturn bool
	ReturnAt  time.Time `gorm:"type:timestamp"`
}

// Umbrella schema
type Umbrella struct {
	ID     int32  `gorm:"AUTO_INCREMENT"`
	Number string `gorm:"type:varchar(10)"`
}
