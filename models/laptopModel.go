package models

import (
	"time"

	"gorm.io/gorm"
)

type Laptop struct {
	ID          int    			`json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name        string 			`json:"name" gorm:"type:varchar(255);not null"`
	Brand       string 			`json:"brand" gorm:"type:varchar(255);not null"`
	Processor   string 			`json:"processor" gorm:"type:varchar(255);not null"`
	RAM         int    			`json:"ram"`
	Storage     int    			`json:"storage"`
	Display     string 			`json:"display" gorm:"type:varchar(255);not null"`
	Graphics    string 			`json:"graphics" gorm:"type:varchar(255);not null"`
	OS          string 			`json:"os" gorm:"type:varchar(255);not null"`
	Battery     string 			`json:"battery" gorm:"type:varchar(255);not null"`
	Keyboard    string 			`json:"keyboard" gorm:"type:varchar(255);not null"`
	Description string 			`json:"description" gorm:"type:text"`
	Price       int    			`json:"price"`
	CreatedAt   time.Time		`json:"created_at"`
	UpdatedAt   time.Time		`json:"updated_at"`
	DeletedAt   gorm.DeletedAt 	`json:"deleted_at" gorm:"index"`
}