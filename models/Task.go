package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Name    string `gorm:"not null;unique_index"`
	Description     string
	UserId uint
}