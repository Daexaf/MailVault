package dto

import "time"

type MailAccountResponse struct {
	ID uint `json:"id"`

	BranchID uint `json:"branchId"`

	Email string `json:"email"`

	IsActive bool `json:"isActive"`

	LastSyncAt *time.Time `json:"lastSyncAt"`
}
