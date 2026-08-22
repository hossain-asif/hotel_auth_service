package rolepermission

import (
	"context"
	"errors"
	"fmt"

	"github.com/hossain-asif/hotel_auth_service/internal/db/models"
	"github.com/hossain-asif/hotel_auth_service/utils/pg"
	"gorm.io/gorm"
)

type RolePermissionRepository interface {
	GetRolePermissionById(ctx context.Context, id string) (*models.RolePermission, error)
	GetRolePermissionByRoleId(ctx context.Context, roleId string) ([]*models.RolePermission, error)
	AddPermissionToRole(ctx context.Context, rolePermission *models.RolePermission) (string, error)
	RemovePermissionFromRole(ctx context.Context, roleId string, permissionId string) error
	GetAllRolePermissions(ctx context.Context) ([]*models.RolePermission, error)
}

type RolePermissionRepositoryImpl struct {
	db *gorm.DB
}

func NewRolePermissionRepository(_db *gorm.DB) RolePermissionRepository {
	return &RolePermissionRepositoryImpl{
		db: _db,
	}
}

func (rp *RolePermissionRepositoryImpl) GetRolePermissionById(ctx context.Context, id string) (*models.RolePermission, error) {
	rolePermission := &models.RolePermission{}
	err := rp.db.First(rolePermission, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return rolePermission, nil
}

func (rp *RolePermissionRepositoryImpl) GetRolePermissionByRoleId(ctx context.Context, roleId string) ([]*models.RolePermission, error) {
	var rolePermissions []*models.RolePermission
	if err := rp.db.Where("role_id = ?", roleId).Find(&rolePermissions).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return rolePermissions, nil
}

func (rp *RolePermissionRepositoryImpl) AddPermissionToRole(ctx context.Context, rolePermission *models.RolePermission) (string, error) {
	result := rp.db.Create(&rolePermission)

	if result.Error != nil {
		err := pg.HandlePgError(result.Error)
		return "", err
	}

	if result.RowsAffected == 0 {
		return "", fmt.Errorf("no Role Permission was created")
	}
	return fmt.Sprintf("Created Role Permission (rows affected: %d)\n", result.RowsAffected), nil
}

func (rp *RolePermissionRepositoryImpl) RemovePermissionFromRole(ctx context.Context, roleId string, permissionId string) error {
	result := rp.db.
		Where("role_id = ? AND permission_id = ?", roleId, permissionId).
		Delete(&models.RolePermission{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (rp *RolePermissionRepositoryImpl) GetAllRolePermissions(ctx context.Context,) ([]*models.RolePermission, error) {
	var rolePermissions []*models.RolePermission
	if err := rp.db.Find(&rolePermissions).Error; err != nil {
		return nil, err
	}
	return rolePermissions, nil
}
