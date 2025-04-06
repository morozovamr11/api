package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"index"`
	Password string
	Name     string
}

func NewUser(email string, name string) *User {
	user := &User{
		Email: email,
		Name:  name,
	}
	return user
}
