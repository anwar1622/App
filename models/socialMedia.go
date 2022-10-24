package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `json:"name" form:"name" validate:"required~Name is required"`
	SocialMediaURl string `json:"social_media_url" form:"social_media_url" validate:"required~Social Media URL is required"`
	UserID         uint
	User           *User
}

func (p *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return err
	}
	return
}

func (p *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)
	if errUpdate != nil {
		err = errUpdate
		return err
	}
	err = nil
	return
}
