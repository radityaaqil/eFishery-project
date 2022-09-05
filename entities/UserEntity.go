package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// type User struct {
// 	gorm.Model
// 	Username string `gorm:"column:username" json:"username"`
// 	Email    string `gorm:"column:email" json:"email"`
// 	Password string `gorm:"column:password" json:"password"`
// }
