package entity

// import "github.com/jinzhu/gorm"

// import (
// 	"github.com/google/uuid"
// )

type Inventorys struct {
	ID int `gorm:"primary_key;column:uuid;type:int(32) NOT NULL AUTO_INCREMENT" json:"uuid"`
	// UUID         uuid.UUID
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Availability int     `json:"availability"`
	IsActive     bool    `json:"is_active"`
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}
