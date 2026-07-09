package gorm

import (
	"github.com/daexaf/mailvault/internal/domain/entities"
	"github.com/daexaf/mailvault/internal/domain/repository"
	"gorm.io/gorm"
)

type branchRepository struct {
	db *gorm.DB
}

func NewBranchRepository(db *gorm.DB) repository.BranchRepository {
	return &branchRepository{
		db: db,
	}
}

func (r *branchRepository) Create(branch *entities.Branch) error {
	return r.db.Create(branch).Error
}

func (r *branchRepository) FindByID(id uint) (*entities.Branch, error) {
	var branch entities.Branch

	err := r.db.First(&branch, id).Error
	if err != nil {
		return nil, err
	}

	return &branch, nil
}

func (r *branchRepository) FindAll() ([]entities.Branch, error) {
	var branches []entities.Branch

	if err := r.db.Find(&branches).Error; err != nil {
		return nil, err
	}

	return branches, nil
}

func (r *branchRepository) ExistsByCode(code string) (bool, error) {
	var count int64

	err := r.db.
		Model(&entities.Branch{}).
		Where("code = ?", code).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
