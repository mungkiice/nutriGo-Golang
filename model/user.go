package model

import "golang.org/x/crypto/bcrypt"

type User struct{
	Model
	Nama		string		`gorm:"type:varchar(191);not_null"`
	Email		string		`gorm:"type:varchar(191);unique;not_null"`
	Password	string		`gorm:"type:varchar(191);not_null"`
	Gender		string		`gorm:"type:varchar(191)" sql:"type:ENUM('laki-laki','perempuan')"`
	TinggiBadan	int			`gorm:"type:double"`
	BeratBadan	int			`gorm:"type:double"`
	Usia		int			`gorm:"type:int(11)"`
	IsAdmin		bool		`gorm:"type:tinyint(1);default:0"`
	History		[]History
}

func (User) TableName() string {
	return "user"
}

func (u *User) VerifyPassword(pass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass)); err != nil {
		return err
	}
	return nil
}
