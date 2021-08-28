package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Street   string `gorm:"type: varchar(50);not null"`
	Area     string
	Town     string
	Province string
}
