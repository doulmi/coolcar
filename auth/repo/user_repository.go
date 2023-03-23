package repo

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"size:255;index:unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func setup(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

type UserRepository struct {
	// db *
}
