package rolerepo

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

type RoleRepository interface {
	Create(ctx context.Context, role *models.Role) (string, error)
	Update(ctx context.Context, id string, updatePayload *dto.RoleUpadateRequestPayload) (string, error)
	SoftDelete(ctx context.Context, id string) (string, error)
	HardDelete(ctx context.Context, id string) (string, error)

	GetByID(ctx context.Context, id string) (*models.Role, error)
	GetAll(ctx context.Context) ([]*models.Role, error)

	GetRoleByName(ctx context.Context, name string) (*models.Role, error)

}

type RoleRepositoryImpl struct {
	// Add fields for database connection, etc.
	db                *gorm.DB
	RoleRepositoryLog *logger.GormLogWriter
}

func NewRoleRepository(_db *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{
		db:                _db,
		RoleRepositoryLog: &logger.GormLogWriter{Logger: logger.Log.Scope("repository", "gorm", "Role_repository")},
	}
}

func (r *RoleRepositoryImpl) Create(ctx context.Context, role *models.Role) (string, error) {
	log := r.RoleRepositoryLog.Method("Create").WithContext(ctx)

	result := r.db.Create(&role)

	if result.Error != nil {
		err := pg.HandlePgError(result.Error)
		log.Errorf("Error creating Role: %v\n", err)
		return "", err
	}

	if result.RowsAffected == 0 {
		log.Error("No Role was created.")
		return "", fmt.Errorf("no Role was created")
	}

	return fmt.Sprintf("Created Role (rows affected: %d)\n", result.RowsAffected), nil
}

func (r *RoleRepositoryImpl) GetByID(ctx context.Context, id string) (*models.Role, error) {
	log := r.RoleRepositoryLog.Method("GetByID").WithContext(ctx)
	role := &models.Role{}

	err := r.db.
		Select("id", "name", "description", "created_at", "updated_at").
		First(role, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("Role not found.")
			return nil, err
		}
		log.Errorf("Error fetching Role: %v\n", err)
		return nil, err
	}
	return role, nil
}

func (r *RoleRepositoryImpl) GetAll(ctx context.Context) ([]*models.Role, error) {
	log := r.RoleRepositoryLog.Method("GetAll").WithContext(ctx)

	var roles []*models.Role

	err := r.db.
		Select("id", "name", "description", "created_at", "updated_at").
		Find(&roles).Error

	if err != nil {
		log.Errorf("Error fetching Roles: %v\n", err)
		return nil, err
	}

	return roles, nil
}

func (r *RoleRepositoryImpl) Update(ctx context.Context, id string, updatePayload *dto.RoleUpadateRequestPayload) (string, error) {
	log := r.RoleRepositoryLog.Method("Update").WithContext(ctx)

	fields := map[string]interface{}{}
	if updatePayload.Name != "" {
		fields["name"] = updatePayload.Name
	}
	if updatePayload.Description != nil {
		fields["description"] = *updatePayload.Description
	}
	if len(fields) == 0 {
		log.Errorf("No fields to update.")
		return "", fmt.Errorf("no fields to update")
	}

	result := r.db.
		Model(&models.Role{}).
		Where("id = ?", id).
		Updates(fields)

	if result.Error != nil {
		log.Errorf("Error updating Role: %v\n", result.Error)
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		log.Errorf("No Role was updated.")
		return "", fmt.Errorf("no Role was updated")
	}

	return fmt.Sprintf("Role updated successfully (rows affected: %d)", result.RowsAffected), nil
}

func (r *RoleRepositoryImpl) SoftDelete(ctx context.Context, id string) (string, error) {
	log := r.RoleRepositoryLog.Method("SoftDelete").WithContext(ctx)

	result := r.db.
		Where("id = ?", id).
		Delete(&models.Role{})

	if result.Error != nil {
		log.Errorf("Error deleting Role: %v\n", result.Error)
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		log.Errorf("No Role was deleted.")
		return "", fmt.Errorf("no Role was deleted")
	}

	return fmt.Sprintf("Deleted Role (rows affected: %d)\n", result.RowsAffected), nil
}

func (r *RoleRepositoryImpl) HardDelete(ctx context.Context, id string) (string, error) {
	log := r.RoleRepositoryLog.Method("HardDelete").WithContext(ctx)

	result := r.db.
		Unscoped().
		Where("id = ?", id).
		Delete(&models.Role{})

	if result.Error != nil {
		log.Errorf("Error deleting Role: %v\n", result.Error)
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		log.Errorf("No Role was deleted.")
		return "", fmt.Errorf("no Role was deleted")
	}

	return fmt.Sprintf("Deleted Role (rows affected: %d)\n", result.RowsAffected), nil
}


func (r *RoleRepositoryImpl) GetRoleByName(ctx context.Context, name string) (*models.Role, error) {
	log := r.RoleRepositoryLog.Method("GetRoleByName").WithContext(ctx)
	role := &models.Role{}

	err := r.db.
		Select("id", "name", "description", "created_at", "updated_at").
		First(role, "name = ?", name).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("Role not found.")
			return nil, err
		}
		log.Errorf("Error fetching Role: %v\n", err)
		return nil, err
	}
	return role, nil
}