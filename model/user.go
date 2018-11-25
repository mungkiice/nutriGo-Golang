package model

type User struct{
	Model
	Nama		string		`gorm:"type:varchar(191);not_null"`
	Email		string		`gorm:"type:varchar(191);unique;not_null"`
	Password	string		`gorm:"type:varchar(191);not_null"`
	Gender		string		`gorm:"type:varchar(191)" sql:"type:ENUM('laki-laki','perempuan')"`
	TinggiBadan	int			`gorm:"type:double"`
	BeratBadan	int			`gorm:"type:double"`
	Usia		int			`gorm:"type:int(11)"`
	IsAdmin		bool		`gorm:"type:tinyint(1)"`
	History		[]History
}
