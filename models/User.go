package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not null" json:"firstName"`
	LastName  string `gorm:"not null" json:"lastName"`
	Email     string `gorm:"type:varchar(100);unique_index" json:"email"`
	Tasks     []Task `json:"tasks"`
}
