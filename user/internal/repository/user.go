package repository

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"user/internal/service"
)

type User struct {
	UserId         uint   `gorm:"primarykey"`
	Username       string `gorm:"unique"`
	NickName       string
	PasswordDigest string //密文
}

const (
	PasswordCost = 12 // 密码加密难度
)

// CheckUserExist checks if user exists
func (user *User) CheckUserExist(req *service.UserRequest) bool {
	if err := DB.Where("user_name = ?", req.UserName).First(&user).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

// ShowUserInfo Get UserInfo
func (user *User) ShowUserInfo(req *service.UserRequest) (err error) {
	if exist := user.CheckUserExist(req); exist {
		return nil
	}
	return errors.New("UserName Not Exist")
}

func (*User) UserCreate(req *service.UserRequest) error {
	var count int64
	DB.Where("user_name=?", req.UserName).Count(&count)
	if count != 0 {
		return errors.New("UserName Exists")
	}
	user := User{
		Username: req.UserName,
		NickName: req.NickName,
	}

	// password digest
	_ = user.SetPassword(req.Password)
	err := DB.Create(&user).Error
	return err

}

// SetPassword password encryption
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword decrypt password
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

// BuildUser builds user info
func BuildUser(item User) *service.UserModel {
	userModel := service.UserModel{
		UserID:   uint32(item.UserId),
		UserName: item.NickName,
		NickName: item.NickName,
	}
	return &userModel
}
