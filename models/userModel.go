package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int    			`json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name      string 			`json:"name" gorm:"type:varchar(255);not null"`
	Email     string 			`json:"email" gorm:"type:varchar(255);not null;unique"`
	Password  string 			`json:"-" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time			`json:"created_at"`
	UpdatedAt time.Time			`json:"updated_at"`
	DeletedAt gorm.DeletedAt 	`json:"deleted_at" gorm:"index"`
}