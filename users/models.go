package users

import (
	"errors"
	"github.com/Panmax/chaos-study-api/common"
	"github.com/Panmax/chaos-study-api/plans"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	common.Model

	Username     string `gorm:"size:64;unique_index;not null"`
	Bio          string `gorm:"size:1024;not null"`
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

func CreateUser(user *UserModel) error {
	tx := common.GetDB().Begin()

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&plans.PlanModel{UserId: user.ID, Count: 1, NotRepeat: true}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil

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
