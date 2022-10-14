package models

type Photo struct {
	Id        int    `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	PhotoUrl  string `json:"photo_url"`
	UserId    int    `json:"user_id" gorm:"foreign_key"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
