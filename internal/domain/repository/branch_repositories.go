package repository

import (
	"github.com/daexaf/mailvault/internal/domain/entities"
)

type BranchRepository interface {
	Create(branch *entities.Branch) error

	FindByID(id uint) (*entities.Branch, error)

	FindAll() ([]entities.Branch, error)

	ExistsByCode(code string) (bool, error)
}
