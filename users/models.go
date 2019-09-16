package users

import (
	"errors"
	"github.com/Panmax/chaos-study-api/common"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	gorm.Model
	Username     string `gorm:"column:username;unique_index"`
	Bio          string `gorm:"column:bio;size:1024"`
	PasswordHash string `gorm:"column:password;not null"`
}

func (UserModel) TableName() string {
	return "user"
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func (u *UserModel) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty!")
	}
	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

func (u *UserModel) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func FindOneUser(username string) (UserModel, error) {
	db := common.GetDB()
	var model UserModel
	err := db.Where("username = ?", username).First(&model).Error
	return model, err
}
