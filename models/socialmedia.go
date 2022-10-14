package models

type SocialMedia struct {
	Id             int    `json:"id" gorm:"primary_key"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId         int    `json:"user_id" gorm:"foreign_key"`
}
