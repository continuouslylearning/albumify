package users

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username" json:"username" binding:"required"`
	Password string `gorm:"column:password" json:"password" binding:"required"`
}

func (u *User) Normalize() map[string]interface{} {
	return map[string]interface{}{
		"ID":       u.ID,
		"username": u.Username,
	}
}
