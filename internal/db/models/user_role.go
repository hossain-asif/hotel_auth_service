package models

import "gorm.io/gorm"

// user_roles — 20260820215418_create_user_roles_table.sql
// FKs: 20260821113401 (user_id → users.id), 20260821113448 (role_id → roles.id)
type UserRole struct {
	gorm.Model
	UserID uint `gorm:"not null"`
	RoleID uint `gorm:"not null"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Role Role `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE"`
}