package dto

type CreateMailAccountRequest struct {
	BranchID uint   `json:"branchId" binding:"required"`
	Provider string `json:"provider" binding:"required"`

	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
