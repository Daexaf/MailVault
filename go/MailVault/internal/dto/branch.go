package dto

type CreateBranchRequest struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
