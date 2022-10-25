package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/rama4zis/go-mygram-aplication/helpers"
	"gorm.io/gorm"
)

type User struct {
	Id        uint       `json:"id" gorm:"primary_key"`
	Username  string     `json:"username" gorm:"unique;not null"`
	Email     string     `json:"email" gorm:"unique;not null"`
	Password  string     `json:"password" gorm:"not null"`
	Age       int        `json:"age" gorm:"not null"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		return errCreate
	}

	u.Password = helpers.HashPassword(u.Password)
	err = nil
	return

}
