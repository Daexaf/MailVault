package handler

import (
	"errors"
	"net/http"

	"github.com/daexaf/mailvault/internal/dto"
	"github.com/daexaf/mailvault/internal/service"
	"github.com/gin-gonic/gin"
)

type MailAccountHandler struct {
	service *service.MailAccountService
}

func NewMailAccountHandler(service *service.MailAccountService) *MailAccountHandler {
	return &MailAccountHandler{
		service: service,
	}
}

func (h *MailAccountHandler) Create(c *gin.Context) {
	var req dto.CreateMailAccountRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	account, err := h.service.Create(req)

	if err != nil {
		switch {
		case errors.Is(err, service.ErrMailAccountAlreadyExists):
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})

		case errors.Is(err, service.ErrUnsupportedProvider):
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

		case errors.Is(err, service.ErrBranchNotFound):
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
		}

		return
	}

	response := dto.MailAccountResponse{
		ID:         account.ID,
		BranchID:   account.BranchID,
		Email:      account.Email,
		IsActive:   account.IsActive,
		LastSyncAt: account.LastSyncAt,
	}

	c.JSON(http.StatusCreated, response)
}
