package models

import (
	"gorm.io/gorm"
)

// unique constraint : uq_permissions_name
type Permission struct {
	gorm.Model
	Name        string  `gorm:"column:name;size:255;uniqueIndex:uq_permissions_name;not null"`
	Description *string `gorm:"column:description;type:text"`
	Resource    string  `gorm:"column:resource;size:100;not null"` // "booking", "room", "user"
	Action      string  `gorm:"column:action;size:100;not null"`   // "create", "read", "update", "delete"

	RolePermissions []RolePermission `gorm:"foreignKey:PermissionID"`
}

func (Permission) TableName() string { return "permissions" }
