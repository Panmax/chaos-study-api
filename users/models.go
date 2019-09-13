package users

import "github.com/jinzhu/gorm"

type UserModel struct {
	gorm.Model
	Username     string  `gorm:"column:username"`
	Email        string  `gorm:"column:email;unique_index"`
	Bio          string  `gorm:"column:bio;size:1024"`
	Image        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password;not null"`
}

func (UserModel) TableName() string {
	return "user"
}
