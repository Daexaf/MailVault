package service

import (
	"errors"

	"github.com/daexaf/mailvault/internal/domain/entities"
	"github.com/daexaf/mailvault/internal/domain/repository"
	"github.com/daexaf/mailvault/internal/dto"
)

type BranchService struct {
	repo repository.BranchRepository
}

func NewBranchService(repo repository.BranchRepository) *BranchService {
	return &BranchService{
		repo: repo,
	}
}

func (s *BranchService) Create(req dto.CreateBranchRequest) (*entities.Branch, error) {
	exists, err := s.repo.ExistsByCode(req.Code)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("branch code already exists")
	}

	branch := entities.Branch{
		Code: req.Code,
		Name: req.Name,
	}

	// return s.repo.Create(&branch)
	if err := s.repo.Create(&branch); err != nil {
		return nil, err
	}

	return &branch, nil
}
