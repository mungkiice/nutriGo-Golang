package model

type History struct{
	Model
	UserID		uint
	User		User
	PolaID		uint
	Pola		Pola
	TinggiBadan	float64	`gorm:"type:double"`
	BeratBadan	float64	`gorm:"type:double"`
}

func (History) TableName() string{
	return "history"
}