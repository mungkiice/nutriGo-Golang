package model

import "golang.org/x/crypto/bcrypt"

type User struct{
	Model
	Nama		string		`gorm:"type:varchar(191);not_null"`
	Email		string		`gorm:"type:varchar(191);unique;not_null"`
	Password	string		`gorm:"type:varchar(191);not_null"`
	Gender		string		`gorm:"type:varchar(191)" sql:"type:ENUM('laki-laki','perempuan')"`
	TinggiBadan	float64		`gorm:"type:double"`
	BeratBadan	float64		`gorm:"type:double"`
	Usia		int			`gorm:"type:int(11)"`
	IsAdmin		bool		`gorm:"type:tinyint(1);default:0"`
	Histories	[]History
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
