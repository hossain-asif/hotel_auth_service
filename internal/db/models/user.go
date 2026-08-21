package models

import (
	"time"

	"gorm.io/gorm"
)

// users — 20260212174735_create_user_table.sql
type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255;uniqueIndex:idx_users_email;not null"`
	Password string `gorm:"size:255;not null"`

	// join rows from user_roles
	UserRoles []UserRole `gorm:"foreignKey:UserID"`
}

// Cursorable — used by your pagination package
func (u User) GetID() uint             { return u.ID }
func (u User) GetCreatedAt() time.Time { return u.CreatedAt }