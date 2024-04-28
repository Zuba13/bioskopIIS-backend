package model

type Director struct {
	ID        uint    `gorm:"primary_key" json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Movies    []Movie `gorm:"many2many:director_movies;" json:"movies"`
}
