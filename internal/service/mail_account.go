// package service

// import (
// 	"errors"
// 	"fmt"
// 	"strings"
// 	"time"

// 	"github.com/daexaf/mailvault/internal/domain/entities"
// 	"github.com/daexaf/mailvault/internal/domain/repository"
// 	"github.com/daexaf/mailvault/internal/dto"
// 	"github.com/daexaf/mailvault/internal/imap"
// 	"github.com/daexaf/mailvault/storage"

// 	"gorm.io/gorm"
// )

// var ErrMailAccountAlreadyExists = errors.New("Mail Account already exists")

// var ErrUnsupportedProvider = errors.New("Unsupported Provider")

// var ErrMailAccountNotFound = errors.New("Email Account was Not Found")

// type MailAccountService struct {
// 	repo           repository.MailAccountRepository
// 	branchRepo     repository.BranchRepository
// 	emailRepo      repository.EmailRepository
// 	attachmentRepo repository.AttachmentRepository
// }

// func NewMailAccountService(
// 	repo repository.MailAccountRepository,
// 	branchRepo repository.BranchRepository,
// 	emailRepo repository.EmailRepository,
// 	attachmentRepo repository.AttachmentRepository,
// ) *MailAccountService {
// 	return &MailAccountService{
// 		repo:           repo,
// 		branchRepo:     branchRepo,
// 		emailRepo:      emailRepo,
// 		attachmentRepo: attachmentRepo,
// 	}
// }

// type SyncResult struct {
// 	Fetched int
// 	Saved   int
// 	Skipped int
// }

// func (s *MailAccountService) Create(req dto.CreateMailAccountRequest) (*entities.MailAccount, error) {
// 	_, err := s.branchRepo.FindByID(req.BranchID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	exists, err := s.repo.ExistsByEmail(req.Email)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if exists {
// 		return nil, ErrMailAccountAlreadyExists
// 	}

// 	var (
// 		host string
// 		port int
// 		ssl  bool
// 	)

// 	switch strings.ToLower(req.Provider) {
// 	case "gmail":
// 		host = "imap.gmail.com"
// 		port = 993
// 		ssl = true

// 	case "yahoo":
// 		host = "imap.mail.yahoo.com"
// 		port = 993
// 		ssl = true

// 	default:
// 		return nil, ErrUnsupportedProvider
// 	}

// 	account := entities.MailAccount{
// 		BranchID: req.BranchID,
// 		Email:    req.Email,
// 		Password: req.Password,

// 		Provider: strings.ToLower(req.Provider),
// 		Host:     host,
// 		Port:     port,
// 		UseSSL:   ssl,
// 	}

// 	if err := s.repo.Create(&account); err != nil {
// 		return nil, err
// 	}

// 	return &account, nil
// }

// // func (s *MailAccountService) Sync(id uint) (*SyncResult, error) {
// // 	account, err := s.repo.FindByID(id)
// // 	if err != nil {
// // 		if errors.Is(err, gorm.ErrRecordNotFound) {
// // 			return nil, ErrMailAccountNotFound
// // 		}
// // 		return nil, err
// // 	}
// // 	messages, err := imap.FetchMessage(
// // 		account.Host,
// // 		account.Port,
// // 		account.Email,
// // 		account.Password,
// // 	)
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	result := &SyncResult{
// // 		Fetched: len(messages),
// // 	}

// // 	for _, message := range messages {
// // 		//message-id kosong tidak bagus dijadikan identified duplicate
// // 		if message.MessageID == "" {
// // 			continue
// // 		}

// // 		exists, err := s.emailRepo.ExistByProviderMessageID(
// // 			message.MessageID,
// // 		)

// // 		if err != nil {
// // 			return nil, err
// // 		}
// // 		if exists {
// // 			result.Skipped++
// // 			continue
// // 		}

// // 		emailEntity := entities.Email{
// // 			MailAccountID:     account.ID,
// // 			ProviderMessageID: message.MessageID,
// // 			To:                message.To,
// // 			From:              message.From,
// // 			Subject:           message.Subject,
// // 			Body:              message.Body,
// // 			CC:                message.CC,
// // 			BCC:               message.BCC,
// // 			ReceivedAt:        message.Date,
// // 			HasAttachment:     message.HasAttachment,
// // 			Folder:            "INBOX",
// // 		}
// // 		if err := s.emailRepo.Create(&emailEntity); err != nil {
// // 			return nil, err
// // 		}

// // 		for _, attachment := range message.Attachments {
// // 			storedName, storagePath, err := storage.SaveAttachment(
// // 				"storage",
// // 				account.ID,
// // 				emailEntity.ID,
// // 				attachment.FileName,
// // 				attachment.Data,
// // 			)
// // 			if err != nil {
// // 				_ = s.emailRepo.DeletePermanently(&emailEntity)
// // 				return nil, err
// // 			}

// // 			attachmentEntity := entities.Attachment{
// // 				EmailID:      emailEntity.ID,
// // 				OriginalName: attachment.FileName,
// // 				StoredName:   storedName,
// // 				StoragePath:  storagePath,
// // 				FileSize:     int64(len(attachment.Data)),
// // 				MimeType:     attachment.ContentType,
// // 			}

// // 			if err := s.attachmentRepo.Create(&attachmentEntity); err != nil {
// // 				_ = s.emailRepo.DeletePermanently(&emailEntity)
// // 				return nil, err
// // 			}
// // 		}

// // 		result.Saved++
// // 	}
// // 	return result, nil
// // }

// func (s *MailAccountService) TestConnection(id uint) error {
// 	account, err := s.repo.FindByID(id)
// 	if err != nil {

// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return ErrMailAccountNotFound
// 		}

// 		return err
// 	}

// 	messages, err := imap.FetchMessage(
// 		account.Host,
// 		account.Port,
// 		account.Email,
// 		account.Password,
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	location, err := time.LoadLocation("Asia/Jakarta")
// 	if err != nil {
// 		return err
// 	}

// 	// return imap.TestConnection(
// 	// 	account.Host,
// 	// 	account.Port,
// 	// 	account.Email,
// 	// 	account.Password,
// 	// )
// 	// fmt.Println("===== LATEST EMAIL =====")
// 	// fmt.Println("Message ID :", message.MessageID)
// 	// fmt.Println("Subject    :", message.Subject)
// 	// fmt.Println("From       :", message.From)
// 	// fmt.Println("Date       :", message.Date.In(location).Format("2006-01-02 15:04:05"))
// 	// fmt.Println("========================")

// 	// ini semua email yang diambil dari inbox, bisa lebih dari satu
// 	// fmt.Println("===== EMAILS =====")

// 	// for _, message := range messages {
// 	// 	fmt.Println("------------------------")
// 	// 	fmt.Println("Message ID :", message.MessageID)
// 	// 	fmt.Println("Subject    :", message.Subject)
// 	// 	fmt.Println("From       :", message.From)
// 	// 	fmt.Println("Date       :", message.Date.In(location).Format("2006-01-02 15:04:05"))
// 	// }

// 	// fmt.Println("==================")

// 	fmt.Println("===== LAST 5 EMAILS =====")

// 	start := len(messages) - 5

// 	if start < 0 {
// 		start = 0
// 	}

// 	for i := start; i < len(messages); i++ {
// 		message := messages[i]

// 		fmt.Println("------------------------")
// 		fmt.Println("Message ID :", message.MessageID)
// 		fmt.Println("Subject    :", message.Subject)
// 		fmt.Println("From       :", message.From)
// 		fmt.Println("To         :", message.To)
// 		fmt.Println("CC         :", message.CC)
// 		fmt.Println("BCC        :", message.BCC)
// 		fmt.Println("Date       :", message.Date.In(location).Format("2006-01-02 15:04:05"))
// 		fmt.Println("Body Length :", len(message.Body))
// 		fmt.Println("Has Attachment :", message.HasAttachment)
// 		fmt.Println("Attachment Count:", len(message.Attachments))

// 		for _, attachment := range message.Attachments {
// 			fmt.Println("  File Name    :", attachment.FileName)
// 			fmt.Println("  Content Type :", attachment.ContentType)
// 			fmt.Println("  File Size    :", len(attachment.Data))
// 		}
// 	}

// 	fmt.Println("=========================")

// 	return nil
// }

// // func (s *MailAccountService) Sync(id uint) error {

// // }

package service

import (
	"errors"
	"strings"

	"github.com/daexaf/mailvault/internal/domain/entities"
	"github.com/daexaf/mailvault/internal/domain/repository"
	"github.com/daexaf/mailvault/internal/dto"
	"github.com/daexaf/mailvault/internal/imap"
	"github.com/daexaf/mailvault/storage"

	"gorm.io/gorm"
)

var ErrMailAccountAlreadyExists = errors.New("Mail Account already exists")

var ErrUnsupportedProvider = errors.New("Unsupported Provider")

var ErrMailAccountNotFound = errors.New("Email Account was Not Found")

type MailAccountService struct {
	repo           repository.MailAccountRepository
	branchRepo     repository.BranchRepository
	emailRepo      repository.EmailRepository
	attachmentRepo repository.AttachmentRepository
}

func NewMailAccountService(
	repo repository.MailAccountRepository,
	branchRepo repository.BranchRepository,
	emailRepo repository.EmailRepository,
	attachmentRepo repository.AttachmentRepository,
) *MailAccountService {
	return &MailAccountService{
		repo:           repo,
		branchRepo:     branchRepo,
		emailRepo:      emailRepo,
		attachmentRepo: attachmentRepo,
	}
}

type SyncResult struct {
	Fetched int
	Saved   int
	Skipped int
}

func (s *MailAccountService) Create(
	req dto.CreateMailAccountRequest,
) (*entities.MailAccount, error) {

	// Pastikan Branch ada
	_, err := s.branchRepo.FindByID(req.BranchID)
	if err != nil {
		return nil, err
	}

	// Cek apakah email sudah terdaftar
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

	// Tentukan konfigurasi berdasarkan provider
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

		Provider: strings.ToLower(req.Provider),
		Host:     host,
		Port:     port,
		UseSSL:   ssl,
	}

	if err := s.repo.Create(&account); err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *MailAccountService) Sync(id uint) (*SyncResult, error) {

	// =========================================
	// 1. Ambil Mail Account
	// =========================================

	account, err := s.repo.FindByID(id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMailAccountNotFound
		}

		return nil, err
	}

	const batchSize uint32 = 100

	// =========================================
	// 2. Ambil metadata seluruh email
	//
	// Belum download body dan attachment.
	// Jadi proses ini relatif ringan.
	// =========================================

	metadataList, err := imap.FetchMessageMetadataByUID(
		account.Host,
		account.Port,
		account.Email,
		account.Password,
		account.LastUID,
		batchSize,
	)

	if err != nil {
		return nil, err
	}

	result := &SyncResult{
		Fetched: len(metadataList),
	}

	// Tidak ada email baru.
	if len(metadataList) == 0 {
		return result, nil
	}

	// =========================================
	// 3. Periksa setiap metadata email
	// =========================================

	for _, metadata := range metadataList {

		// Message-ID kosong tidak bisa digunakan
		// sebagai identifier duplicate.
		if metadata.MessageID == "" {

			// Kita tetap maju berdasarkan UID,
			// supaya sync tidak berhenti selamanya
			// di email yang tidak punya Message-ID.

			if err := s.repo.UpdateLastUID(
				account.ID,
				metadata.UID,
			); err != nil {
				return nil, err
			}

			account.LastUID = metadata.UID

			result.Skipped++

			continue
		}

		// =====================================
		// 4. Cek apakah email sudah ada di DB
		// =====================================

		exists, err := s.emailRepo.ExistByProviderMessageID(
			metadata.MessageID,
		)

		if err != nil {
			return nil, err
		}

		if exists {

			result.Skipped++

			// Walaupun duplicate, UID ini sudah aman
			// untuk dilewati.
			if err := s.repo.UpdateLastUID(
				account.ID,
				metadata.UID,
			); err != nil {
				return nil, err
			}

			account.LastUID = metadata.UID

			continue
		}

		// =====================================
		// 5. Email belum ada.
		//
		// Baru sekarang download full email
		// berdasarkan SequenceNumber.
		// =====================================

		message, err := imap.FetchMessageByUID(
			account.Host,
			account.Port,
			account.Email,
			account.Password,
			metadata.UID,
		)

		if err != nil {
			return nil, err
		}

		// =====================================
		// 6. Buat Email Entity
		// =====================================

		emailEntity := entities.Email{
			MailAccountID:     account.ID,
			ProviderMessageID: message.MessageID,
			To:                message.To,
			From:              message.From,
			Subject:           message.Subject,
			Body:              message.Body,
			CC:                message.CC,
			BCC:               message.BCC,
			ReceivedAt:        message.Date,
			HasAttachment:     message.HasAttachment,
			Folder:            "INBOX",
		}

		// =====================================
		// 7. Simpan email ke database
		// =====================================

		if err := s.emailRepo.Create(&emailEntity); err != nil {
			return nil, err
		}

		// =====================================
		// 8. Simpan attachment
		// =====================================

		for _, attachment := range message.Attachments {

			// Simpan file attachment ke local storage
			storedName, storagePath, err := storage.SaveAttachment(
				"storage",
				account.ID,
				emailEntity.ID,
				attachment.FileName,
				attachment.Data,
			)

			if err != nil {

				// Email sudah terlanjur dibuat.
				// Hapus supaya sync berikutnya
				// bisa mencoba lagi.
				_ = s.emailRepo.DeletePermanently(
					&emailEntity,
				)

				return nil, err
			}

			// =================================
			// 9. Buat metadata Attachment
			// =================================

			attachmentEntity := entities.Attachment{
				EmailID:      emailEntity.ID,
				OriginalName: attachment.FileName,
				StoredName:   storedName,
				StoragePath:  storagePath,
				FileSize:     int64(len(attachment.Data)),
				MimeType:     attachment.ContentType,
			}

			// =================================
			// 10. Simpan attachment ke DB
			// =================================

			if err := s.attachmentRepo.Create(
				&attachmentEntity,
			); err != nil {

				_ = s.emailRepo.DeletePermanently(
					&emailEntity,
				)

				return nil, err
			}
		}

		if err := s.repo.UpdateLastUID(
			account.ID,
			metadata.UID,
		); err != nil {
			return nil, err
		}

		account.LastUID = metadata.UID

		result.Saved++
	}

	return result, nil
}

func (s *MailAccountService) TestConnection(id uint) error {

	// Ambil account dari database
	account, err := s.repo.FindByID(id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrMailAccountNotFound
		}

		return err
	}

	// TestConnection sekarang HANYA mengecek
	// apakah koneksi IMAP berhasil.
	return imap.TestConnection(
		account.Host,
		account.Port,
		account.Email,
		account.Password,
	)
}
