package models

import "time"

type SocialMedia struct {
	Id             uint       `json:"id" gorm:"primary_key"`
	Name           string     `json:"name" gorm:"not null"`
	SocialMediaUrl string     `json:"social_media_url" gorm:"not null"`
	UserId         uint       `json:"user_id" gorm:"foreign_key"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}
