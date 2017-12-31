package model

import (
	"time"
)

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type Message struct {
	Model
	TwitterId string `gorm:"type:text;" json:"twitter_id"`
	Body      string `json:"body"`
}

type Profile struct {
	Icon     string `json:"icon"`
	NickName string `json:"nickname"`
}
