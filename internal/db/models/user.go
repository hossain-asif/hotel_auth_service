package models

import (
	"time"

	"gorm.io/gorm"
)

// unique constraint : uq_users_email       
// index             : idx_users_deleted_at
type User struct {
	gorm.Model
	Name     string `gorm:"column:name;size:255;not null"`
	Email    string `gorm:"column:email;size:255;uniqueIndex:uq_users_email;not null"`
	Password string `gorm:"column:password;size:255;not null"`

	// join rows from user_roles
	UserRoles []UserRole `gorm:"foreignKey:UserID"`
}

func (User) TableName() string { return "users" }

// seek_pagination.Entity
func (u User) GetID() uint             { return u.ID }
func (u User) GetCreatedAt() time.Time { return u.CreatedAt }