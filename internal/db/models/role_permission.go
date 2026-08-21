package models

import (
	"gorm.io/gorm"
)

// fk_role_permissions_role_id       -> roles.id       ON DELETE CASCADE
// fk_role_permissions_permission_id -> permissions.id ON DELETE CASCADE
type RolePermission struct {
	gorm.Model
	RoleID       uint `gorm:"column:role_id;not null"`
	PermissionID uint `gorm:"column:permission_id;not null"`

	Role       Role       `gorm:"foreignKey:RoleID;references:ID;constraint:OnDelete:CASCADE"`
	Permission Permission `gorm:"foreignKey:PermissionID;references:ID;constraint:OnDelete:CASCADE"`
}

func (RolePermission) TableName() string { return "role_permissions" }
