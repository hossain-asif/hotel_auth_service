package models

import "gorm.io/gorm"

// role_permissions — 20260820215404_create_role_permissions_table.sql
// FKs: 20260820221415 (role_id → roles.id), 20260821112259 (permission_id → permissions.id)
type RolePermission struct {
	gorm.Model
	RoleID       uint `gorm:"not null"`
	PermissionID uint `gorm:"not null"`

	Role       Role       `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE"`
	Permission Permission `gorm:"foreignKey:PermissionID;constraint:OnDelete:CASCADE"`
}