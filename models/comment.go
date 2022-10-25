package models

import "time"

type Comment struct {
	Id        uint       `json:"id" gorm:"primary_key"`
	UserId    uint       `json:"user_id" gorm:"foreign_key"`
	PhotoId   uint       `json:"photo_id" gorm:"foreign_key"`
	Message   string     `json:"message" gorm:"not null"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
