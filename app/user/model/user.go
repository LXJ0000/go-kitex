package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex;type:varchar(255) not null"`
	Password string `gorm:"type:varchar(255) not null"`
}

func (u User) TableName() string {
	return "t_user"
}

func CreateUser(db *gorm.DB, user *User) (err error) {
	return db.Create(user).Error
}

func GetUserByEmail(db *gorm.DB, email string) (item User, err error) {
	err = db.Where("email = ?", email).First(&item).Error
	return
}
