package handler

import (
	"net/http"

	"github.com/daexaf/mailvault/internal/dto"
	"github.com/daexaf/mailvault/internal/service"
	"github.com/gin-gonic/gin"
)

type BranchHandler struct {
	service *service.BranchService
}

// func (h *BranchHandler) Create(c *gin.Context) {
// 	// Implementation for branch handler

// }

func NewBranchHandler(service *service.BranchService) *BranchHandler {
	return &BranchHandler{
		service: service,
	}
}

func (h *BranchHandler) Create(c *gin.Context) {

	var req dto.CreateBranchRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	branch, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := dto.BranchResponse{
		Code: branch.Code,
		Name: branch.Name,
	}

	c.JSON(http.StatusCreated, response)
}
