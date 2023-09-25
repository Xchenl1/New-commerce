package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `grom:"unique"`
	Email    string
	Password string
	NickName string
	Status   string //状态
	Avatar   string //
	Money    string //余额
}

const (
	PasswordCost        = 12       //加密难度
	Active       string = "active" //已激活
)

func (user *User) TableName() string {
	return "user"
}

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) //如果为空 说明输入的密码和数据库密码一致
	return err == nil
}
