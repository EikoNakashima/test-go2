package model

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Text   string
	Status string
}
