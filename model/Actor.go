package model

type Actor struct {
	ID        uint    `gorm:"primary_key" json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Movies    []Movie `gorm:"many2many:actor_movies;" json:"movies"`
}
