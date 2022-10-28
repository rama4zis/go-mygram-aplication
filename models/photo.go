package models

import "time"

type Photo struct {
	Id        uint       `json:"id" gorm:"primary_key"`
	Title     string     `json:"title" gorm:"not null"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url" gorm:"not null"`
	UserId    uint       `json:"user_id" gorm:"not null"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UserPhoto User       `json:"user" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
}
