package models

import (
	"gorm.io/gorm"
)

// fk_user_roles_user_id -> users.id  ON DELETE CASCADE
// fk_user_roles_role_id -> roles.id  ON DELETE CASCADE
type UserRole struct {
	gorm.Model
	UserID uint `gorm:"column:user_id;not null"`
	RoleID uint `gorm:"column:role_id;not null"`

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Role Role `gorm:"foreignKey:RoleID;references:ID;constraint:OnDelete:CASCADE"`
}

func (UserRole) TableName() string { return "user_roles" }
