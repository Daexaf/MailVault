package service

import (
	"errors"
	"strings"

	"github.com/daexaf/mailvault/internal/domain/entities"
	"github.com/daexaf/mailvault/internal/domain/repository"
	"github.com/daexaf/mailvault/internal/dto"
	"github.com/daexaf/mailvault/internal/imap"
	"gorm.io/gorm"
)

var ErrMailAccountAlreadyExists = errors.New("Mail Account already exists")

var ErrUnsupportedProvider = errors.New("Unsupported Provider")

var ErrMailAccountNotFound = errors.New("Email Account was Not Found")

type MailAccountService struct {
	repo       repository.MailAccountRepository
	branchRepo repository.BranchRepository
}

func NewMailAccountService(
	repo repository.MailAccountRepository,
	branchRepo repository.BranchRepository,
) *MailAccountService {
	return &MailAccountService{
		repo:       repo,
		branchRepo: branchRepo,
	}
}

func (s *MailAccountService) Create(req dto.CreateMailAccountRequest) (*entities.MailAccount, error) {
	_, err := s.branchRepo.FindByID(req.BranchID)
	if err != nil {
		return nil, err
	}

	exists, err := s.repo.ExistsByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, ErrMailAccountAlreadyExists
	}

	var (
		host string
		port int
		ssl  bool
	)

	switch strings.ToLower(req.Provider) {
	case "gmail":
		host = "imap.gmail.com"
		port = 993
		ssl = true

	case "yahoo":
		host = "imap.mail.yahoo.com"
		port = 993
		ssl = true

	default:
		return nil, ErrUnsupportedProvider
	}

	account := entities.MailAccount{
		BranchID: req.BranchID,
		Email:    req.Email,
		Password: req.Password,

		Host:   host,
		Port:   port,
		UseSSL: ssl,
	}

	if err := s.repo.Create(&account); err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *MailAccountService) TestConnection(id uint) error {
	account, err := s.repo.FindByID(id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrMailAccountNotFound
		}

		return err
	}

	return imap.TestConnection(
		account.Host,
		account.Port,
		account.Email,
		account.Password,
	)
}
