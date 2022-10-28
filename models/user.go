package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/rama4zis/go-mygram-aplication/helpers"
	"gorm.io/gorm"
)

type User struct {
	Id        uint       `json:"-" gorm:"primary_key"`
	Username  string     `json:"username" gorm:"unique;not null"`
	Email     string     `json:"email" gorm:"unique;not null"`
	Password  string     `json:"-" gorm:"not null"`
	Age       uint       `json:"-"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
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
