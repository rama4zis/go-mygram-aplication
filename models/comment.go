package models

import (
	"time"
)

type Comment struct {
	Id        uint       `json:"id" gorm:"primary_key"`
	UserId    uint       `json:"user_id" gorm:"foreign_key"`
	PhotoId   uint       `json:"photo_id" gorm:"foreign_key"`
	Message   string     `json:"message" gorm:"not null"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	// User      datatypes.JSON `json:"user"`
	// Photo     datatypes.JSON `json:"photo"`
}
