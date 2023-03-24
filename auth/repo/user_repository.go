package repo

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model

	Email    string `gorm:"size:255;index:user_email_unique_index,unique"`
	Password string
}

func HandleDBMigration(db *gorm.DB) {
	db.AutoMigrate(&Account{})
}
