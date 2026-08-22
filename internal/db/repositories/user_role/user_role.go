package userrolerepo

import (
	"context"
	"fmt"

	"github.com/hossain-asif/hotel_auth_service/common_pkg/logger"
	"github.com/hossain-asif/hotel_auth_service/internal/db/models"
	"github.com/hossain-asif/hotel_auth_service/utils/pg"
	"gorm.io/gorm"
)

type UserRoleRepository interface {
	GetUserRoles(ctx context.Context, userId string) ([]*models.Role, error)
	AssignRoleToUser(ctx context.Context, userRole *models.UserRole) (string, error)
	RemoveRoleFromUser(ctx context.Context, userId string, roleId string) (string, error)
	GetUserPermissions(ctx context.Context, userId string) ([]*models.Permission, error)
	HasPermission(ctx context.Context, userId string, permissionName string) (bool, error)
	HasRole(ctx context.Context, userId string, roleName string) (bool, error)
	HasAllRoles(ctx context.Context, userId string, roleNames []string) (bool, error)
	HasAnyRole(ctx context.Context, userId string, roleNames []string) (bool, error)
}

type UserRoleRepositoryImpl struct {
	// Add fields for database connection, etc.
	db                *gorm.DB
	RoleRepositoryLog *logger.GormLogWriter
}

func NewRoleRepository(_db *gorm.DB) *UserRoleRepositoryImpl {
	return &UserRoleRepositoryImpl{
		db:                _db,
		RoleRepositoryLog: &logger.GormLogWriter{Logger: logger.Log.Scope("userrolerepo;", "gorm", "User Role_repository")},
	}
}

func (ur *UserRoleRepositoryImpl) GetUserRoles(ctx context.Context, userId string) ([]*models.Role, error) {
	query := `
		SELECT r.id, r.name, r.description, r.created_at, r.updated_at
		FROM user_roles ur
		INNER JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = ?`

	var roles []*models.Role
	if err := ur.db.Raw(query, userId).Scan(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (ur *UserRoleRepositoryImpl) AssignRoleToUser(ctx context.Context, userRole *models.UserRole) (string, error) {
	log := ur.RoleRepositoryLog.Method("AssignRoleToUser").WithContext(ctx)

	result := ur.db.Create(&userRole)

	if result.Error != nil {
		err := pg.HandlePgError(result.Error)
		log.Errorf("Error creating user-role: %v\n", err)
		return "", err
	}

	if result.RowsAffected == 0 {
		log.Error("No user-role was created.")
		return "", fmt.Errorf("no user-role was created")
	}

	return fmt.Sprintf("Created user-role (rows affected: %d)\n", result.RowsAffected), nil
}

func (ur *UserRoleRepositoryImpl) RemoveRoleFromUser(ctx context.Context, userId string, roleId string) (string, error) {
	log := ur.RoleRepositoryLog.Method("RemoveRoleFromUser").WithContext(ctx)

	result := ur.db.
		Where("user_id = ? AND role_id = ?", userId, roleId).
		Delete(&models.UserRole{})

	if result.Error != nil {
		log.Errorf("Error deleting UserRole: %v\n", result.Error)
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		log.Errorf("No UserRole was deleted.")
		return "", fmt.Errorf("no UserRole was deleted")
	}

	return fmt.Sprintf("Deleted UserRole (rows affected: %d)\n", result.RowsAffected), nil
}

func (ur *UserRoleRepositoryImpl) GetUserPermissions(ctx context.Context, userId string) ([]*models.Permission, error) {
	query := `
		SELECT p.id, p.name, p.description, p.resource, p.action, p.created_at, p.updated_at
		FROM user_roles ur
		INNER JOIN role_permissions rp ON ur.role_id = rp.role_id
		INNER JOIN permissions p ON rp.permission_id = p.id
		WHERE ur.user_id = ?`

	var permissions []*models.Permission
	if err := ur.db.Raw(query, userId).Scan(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func (ur *UserRoleRepositoryImpl) HasPermission(ctx context.Context, userId string, permissionName string) (bool, error) {
	query := `
		SELECT COUNT(*) > 0
		FROM user_roles ur
		INNER JOIN role_permissions rp ON ur.role_id = rp.role_id
		INNER JOIN permissions p ON rp.permission_id = p.id
		WHERE ur.user_id = ? AND p.name = ?`

	var exists bool
	if err := ur.db.Raw(query, userId, permissionName).Scan(&exists).Error; err != nil {
		return false, err
	}
	return exists, nil
}

func (ur *UserRoleRepositoryImpl) HasRole(ctx context.Context, userId string, roleName string) (bool, error) {
	query := `
		SELECT COUNT(*) > 0
		FROM user_roles ur
		INNER JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = ? AND r.name = ?`

	var exists bool
	if err := ur.db.Raw(query, userId, roleName).Scan(&exists).Error; err != nil {
		return false, err
	}
	return exists, nil
}

func (ur *UserRoleRepositoryImpl) HasAllRoles(ctx context.Context, userId string, roleNames []string) (bool, error) {
	if len(roleNames) == 0 {
		return true, nil
	}

	query := `
		SELECT COUNT(DISTINCT r.name) = ?
		FROM user_roles ur
		INNER JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = ? AND r.name IN (?)`

	var hasAllRoles bool
	if err := ur.db.Raw(query, len(roleNames), userId, roleNames).Scan(&hasAllRoles).Error; err != nil {
		return false, err
	}
	return hasAllRoles, nil
}

func (ur *UserRoleRepositoryImpl) HasAnyRole(ctx context.Context, userId string, roleNames []string) (bool, error) {
	if len(roleNames) == 0 {
		return true, nil // If no roles are specified, return true
	}

	query := `
		SELECT COUNT(*) > 0
		FROM user_roles ur
		INNER JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = ? AND r.name IN (%s)`

	var hasAnyRole bool
	if err := ur.db.Raw(query, userId, roleNames).Scan(&hasAnyRole).Error; err != nil {
		return false, err
	}
	return hasAnyRole, nil
}
