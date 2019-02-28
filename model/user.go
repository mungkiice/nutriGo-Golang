package model

import "golang.org/x/crypto/bcrypt"

type User struct{
	Model
	Nama		string		`json:"nama" gorm:"type:varchar(191);not_null"`
	Email		string		`json:"" gorm:"type:varchar(191);unique;not_null"`
	Password	string		`json:"-" gorm:"type:varchar(191);not_null"`
	Gender		string		`json:"" gorm:"type:varchar(191)" sql:"type:ENUM('laki-laki','perempuan')"`
	TinggiBadan	float64		`json:"" gorm:"type:double"`
	BeratBadan	float64		`json:"" gorm:"type:double"`
	Usia		int			`json:"" gorm:"type:int(11)"`
	IsAdmin		bool		`json:"-" gorm:"type:tinyint(1);default:0"`
	Histories	[]History	`json:"-"`
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
