package models

type Comment struct {
	Id        int    `json:"id" gorm:"primary_key"`
	UserId    int    `json:"user_id" gorm:"foreign_key"`
	PhotoId   int    `json:"photo_id" gorm:"foreign_key"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
