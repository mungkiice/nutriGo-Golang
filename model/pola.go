package model

type Pola struct{
	Model
	LemakTotal			float64		`gorm:"type:double"`
	ProteinTotal		float64		`gorm:"type:double"`
	KarbohidratTotal	float64		`gorm:"type:double"`
	Makanans			[]Makanan	`gorm:"many2many:pola_makan"`
}

func (Pola) TableName() string{
	return "pola"
}