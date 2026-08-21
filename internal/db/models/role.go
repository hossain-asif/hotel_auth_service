package models

import (
	"gorm.io/gorm"
)

// roles — 20260820211406_create_role_table.sql
type Role struct {
	gorm.Model
	Name        string  `gorm:"size:255;index:idx_roles_name;not null"`
	Description *string `gorm:"type:text"`

	UserRoles       []UserRole       `gorm:"foreignKey:RoleID"`
	RolePermissions []RolePermission `gorm:"foreignKey:RoleID"`
}
