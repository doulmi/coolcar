package repo

import (
	rentalpb "coolcar/rental/api/gen/v1"

	"gorm.io/gorm"
)

type Trip struct {
	gorm.Model

	AccountId string
	CardId    string
	Start     string
	End       string
	Current   string
	Status    rentalpb.TripStatus
}

func HandleDBMigration(db *gorm.DB) {
	db.AutoMigrate(&Trip{})
}
