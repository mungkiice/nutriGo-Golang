package model

type Article struct{
	Model
	Judul		string	`gorm:"type:varchar(191)"`
	Deskripsi	string	`gorm:"type:text"`
	Foto		string	`gorm:"type:text"`
}