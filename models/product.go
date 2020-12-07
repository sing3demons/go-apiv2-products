package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID       uint
	Name     string `gorm:"not null"`
	Desc     string `gorm:"not null"`
	Image    string `gorm:"not null"`
	Category string `gorm:"not null"`
	Price    int64  `gorm:"not null"`
}
