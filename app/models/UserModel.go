package models

import (
	"errors"
	"github.com/gowlfer/gowlfer/internal/utils"
	"github.com/gowlfer/gowlfer/internal/utils/database"
	"gorm.io/gorm"
)

type GowlferUser struct {
	gorm.Model
	UserEmail    string `gorm:"not null;unique_index"`
	UserName     string `gorm:"not null"`
	UserPassword string `gorm:"not null"`
	UserCompany  int64
}

func (u *GowlferUser) CreateUsers(email string, name string, password string) error {

	hashedPass, _ := utils.HashPassword(password)

	u.UserEmail = email
	u.UserName = name
	u.UserPassword = hashedPass

	created := database.DB.Create(&u)

	if created.Error != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (u *GowlferUser) CheckPassword(email string, password string) bool {
	var user GowlferUser

	database.DB.Limit(1).Where("user_email = ?", email).Find(&user)

	return utils.CheckPasswordHash(password, user.UserPassword)
}

func (u *GowlferUser) DoesUserExist(email string) bool {
	var user GowlferUser
	database.DB.Limit(1).Where("user_email = ?", email).Find(&user)

	if user.UserEmail == email {
		return true
	}

	return false
}

func (u *GowlferUser) GetUserByID(id string) GowlferUser {
	var user GowlferUser
	found := database.DB.Limit(1).Where("id = ?", id).Find(&user)

	if found.Error != nil {
		return user
	}

	return user
}

func (u *GowlferUser) GetUser(email string) GowlferUser {
	var user GowlferUser
	found := database.DB.Limit(1).Where("user_email = ?", email).Find(&user)

	if found.Error != nil {
		return user
	}

	return user
}
