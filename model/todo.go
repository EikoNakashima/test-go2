package model

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

// type Todo struct {
// 	ID        uint   `json:"id"`
// 	Text      string `json:text"`
// 	Status    string `json:"status"`
// 	CreatedAt time.Time
// }
