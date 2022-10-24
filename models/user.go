package models

import (
	"App/helper"
	"errors"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"not null;unique_index" json:"username" form:"username" validate:"required~Username is required"`
	Email    string `gorm:"not null;unique_index" json:"email" form:"email" validate:"required~Email is required,email~Email is not valid"`
	Password string `json:"password" form:"password" validate:"required~Password is required,min=6~Password must be at least 6 characters"`
	Age      uint   `json:"age" form:"age" validate:"required~Age is required,"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Age < 8 {
		err = errors.New("Minimum age to register is 8")
		return err
	}
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return err
	}
	u.Password = helper.HashPass(u.Password)
	return
}
