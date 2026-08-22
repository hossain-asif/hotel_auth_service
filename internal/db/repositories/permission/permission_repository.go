package permissionrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/hossain-asif/hotel_auth_service/common_pkg/logger"
	"github.com/hossain-asif/hotel_auth_service/internal/db/models"
	"github.com/hossain-asif/hotel_auth_service/internal/dto"
	"github.com/hossain-asif/hotel_auth_service/utils/pg"

	"gorm.io/gorm"
)

type PermissionRepository interface {
	Create(ctx context.Context, permission *models.Permission) (string, error)
	Update(ctx context.Context, id string, updatePayload *dto.PermissionUpadateRequestPayload) (string, error)
	SoftDelete(ctx context.Context, id string) (string, error)
	HardDelete(ctx context.Context, id string) (string, error)

	GetByID(ctx context.Context, id string) (*models.Permission, error)
	GetAll(ctx context.Context) ([]*models.Permission, error)
	GetPermissionByName(ctx context.Context, name string) (*models.Permission, error)

}

type PermissionRepositoryImpl struct {
	// Add fields for database connection, etc.
	db                *gorm.DB
	PermissionRepositoryLog *logger.GormLogWriter
}

func NewPermissionRepository(_db *gorm.DB) PermissionRepository {
	return &PermissionRepositoryImpl{
		db:                _db,
		PermissionRepositoryLog: &logger.GormLogWriter{Logger: logger.Log.Scope("repository", "gorm", "Permission_repository")},
	}
}

func (p *PermissionRepositoryImpl) Create(ctx context.Context, permission *models.Permission) (string, error) {
	log := p.PermissionRepositoryLog.Method("Create").WithContext(ctx)

	result := p.db.Create(&permission)

	if result.Error != nil {
		err := pg.HandlePgError(result.Error)
		log.Errorf("Error creating Permission: %v\n", err)
		return "", err
	}

	if result.RowsAffected == 0 {
		log.Error("No Permission was created.")
		return "", fmt.Errorf("no Permission was created")
	}

	return fmt.Sprintf("Created Permission (rows affected: %d)\n", result.RowsAffected), nil
}

func (p *PermissionRepositoryImpl) GetByID(ctx context.Context, id string) (*models.Permission, error) {
	log := p.PermissionRepositoryLog.Method("GetByID").WithContext(ctx)
	permission := &models.Permission{}

	err := p.db.
		Select("id", "name", "description", "resource", "action", "created_at", "updated_at").
		First(permission, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("Permission not found.")
			return nil, err
		}
		log.Errorf("Error fetching Permission: %v\n", err)
		return nil, err
	}
	return permission, nil
}

func (p *PermissionRepositoryImpl) GetAll(ctx context.Context) ([]*models.Permission, error) {
	log := p.PermissionRepositoryLog.Method("GetAll").WithContext(ctx)

	var Permissions []*models.Permission

	err := p.db.
		Select("id", "name", "description", "resource", "action", "created_at", "updated_at").
		Find(&Permissions).Error

	if err != nil {
		log.Errorf("Error fetching Permissions: %v\n", err)
		return nil, err
	}

	return Permissions, nil
}

func (p *PermissionRepositoryImpl) Update(ctx context.Context, id string, updatePayload *dto.PermissionUpadateRequestPayload) (string, error) {
	log := p.PermissionRepositoryLog.Method("Update").WithContext(ctx)

	fields := map[string]interface{}{}
	if updatePayload.Name != "" {
		fields["name"] = updatePayload.Name
	}
	if updatePayload.Description != nil {
		fields["description"] = *updatePayload.Description
	}
	if updatePayload.Description != nil {
		fields["resource"] = updatePayload.Resource
	}
	if updatePayload.Description != nil {
		fields["action"] = updatePayload.Action
	}
	if len(fields) == 0 {
		log.Errorf("No fields to update.")
		return "", fmt.Errorf("no fields to update")
	}

	result := p.db.
		Model(&models.Permission{}).
		Where("id = ?", id).
		Updates(fields)

	if result.Error != nil {
		log.Errorf("Error updating Permission: %v\n", result.Error)
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		log.Errorf("No Permission was updated.")
		return "", fmt.Errorf("no Permission was updated")
	}

	return fmt.Sprintf("Permission updated successfully (rows affected: %d)", result.RowsAffected), nil
}

func (p *PermissionRepositoryImpl) SoftDelete(ctx context.Context, id string) (string, error) {
	log := p.PermissionRepositoryLog.Method("SoftDelete").WithContext(ctx)

	result := p.db.
		Where("id = ?", id).
		Delete(&models.Permission{})

	if result.Error != nil {
		log.Errorf("Error deleting Permission: %v\n", result.Error)
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		log.Errorf("No Permission was deleted.")
		return "", fmt.Errorf("no Permission was deleted")
	}

	return fmt.Sprintf("Deleted Permission (rows affected: %d)\n", result.RowsAffected), nil
}

func (p *PermissionRepositoryImpl) HardDelete(ctx context.Context, id string) (string, error) {
	log := p.PermissionRepositoryLog.Method("HardDelete").WithContext(ctx)

	result := p.db.
		Unscoped().
		Where("id = ?", id).
		Delete(&models.Permission{})

	if result.Error != nil {
		log.Errorf("Error deleting Permission: %v\n", result.Error)
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		log.Errorf("No Permission was deleted.")
		return "", fmt.Errorf("no Permission was deleted")
	}

	return fmt.Sprintf("Deleted Permission (rows affected: %d)\n", result.RowsAffected), nil
}

func (r *PermissionRepositoryImpl) GetPermissionByName(ctx context.Context, name string) (*models.Permission, error) {
	log := r.PermissionRepositoryLog.Method("GetPermissionByName").WithContext(ctx)
	permission := &models.Permission{}

	err := r.db.
		Select("id", "name", "description", "resource", "action", "created_at", "updated_at").
		First(permission, "name = ?", name).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("permission not found.")
			return nil, err
		}
		log.Errorf("Error fetching Role: %v\n", err)
		return nil, err
	}
	return permission, nil
}