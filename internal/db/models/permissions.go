package models

import (
	"gorm.io/gorm"
)

// permissions — 20260820215231_create_permissions_table.sql
type Permission struct {
	gorm.Model
	Name        string  `gorm:"size:255;index:idx_permissions_name;not null"`
	Description *string `gorm:"type:text"`
	Resource    string  `gorm:"size:100;not null"` // e.g. "booking", "room", "user"
	Action      string  `gorm:"size:100;not null"` // e.g. "create", "read", "update", "delete"

	RolePermissions []RolePermission `gorm:"foreignKey:PermissionID"`
}
