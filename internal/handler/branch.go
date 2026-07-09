package handler

import (
	"errors"
	"net/http"
	"strconv"

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

func (h *BranchHandler) FindAll(c *gin.Context) {
	branches, err := h.service.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	responses := make([]dto.BranchResponse, 0, len(branches))

	for _, branch := range branches {
		responses = append(responses, dto.BranchResponse{
			Code: branch.Code,
			Name: branch.Name,
		})
	}

	c.JSON(http.StatusOK, responses)
}

func (h *BranchHandler) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid branch id",
		})
		return
	}

	branch, err := h.service.FindByID(uint(id))
	if err != nil {
		if errors.Is(err, service.ErrBranchNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "branch not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch branch",
		})
		return
	}

	response := dto.BranchResponse{
		Code: branch.Code,
		Name: branch.Name,
	}

	c.JSON(http.StatusOK, response)

}
