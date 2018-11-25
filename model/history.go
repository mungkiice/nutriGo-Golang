package model

type History struct{
	Model
	UserID		int
	User		User
	PolaID		int
	Pola		Pola
	TinggiBadan	float64	`gorm:"type:double"`
	BeratBadan	float64	`gorm:"type:double"`
}
