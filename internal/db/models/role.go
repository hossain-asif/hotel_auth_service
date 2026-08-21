package models

import (
	"gorm.io/gorm"
)

// unique constraint : uq_roles_name        
// index             : idx_roles_deleted_at
type Role struct {
	gorm.Model
	Name        string  `gorm:"column:name;size:255;uniqueIndex:uq_roles_name;not null"`
	Description *string `gorm:"column:description;type:text"`

	UserRoles       []UserRole       `gorm:"foreignKey:RoleID"`
	RolePermissions []RolePermission `gorm:"foreignKey:RoleID"`
}

func (Role) TableName() string { return "roles" }
