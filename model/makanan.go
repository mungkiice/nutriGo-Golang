package model

type Makanan struct{
	Model
	Nama		string	`gorm:"type:varchar(191)"`
	Lemak		float64	`gorm:"type:double"`
	Protein		float64	`gorm:"type:double"`
	Karbohidrat	float64	`gorm:"type:double"`
	Pola		[]Pola	`gorm:"many2many:pola_makan"`
}

func (Makanan) TableName() string{
	return "makanan"
}