package entity

import "github.com/jinzhu/gorm"

// import "github.com/jinzhu/gorm"

// import (
// 	"github.com/google/uuid"
// )

type Users struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
