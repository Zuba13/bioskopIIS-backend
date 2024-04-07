package model

import (
	"time"
)

type User struct {
    ID        uint      `gorm:"primary_key" json:"id"`
    Username  string    `gorm:"unique;not null" json:"username"`
    Password  string    `gorm:"not null" json:"password"`
    Email     string    `gorm:"unique;not null" json:"email"`
    Role      string    `gorm:"not null" json:"role"`
    FirstName string    `json:"firstName"`
    LastName  string    `json:"lastName"`
    Money     float64   `json:"money"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
