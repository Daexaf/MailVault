package handler

import (
	"net/http"

	"github.com/daexaf/mailvault/internal/service"

	"github.com/gin-gonic/gin"
)

type EmailHandler struct {
	service *service.EmailService
}

func NewEmailHandler(
	service *service.EmailService,
) *EmailHandler {
	return &EmailHandler{
		service: service,
	}
}

func (h *EmailHandler) Search(c *gin.Context) {

	subject := c.Query("subject")
	from := c.Query("from")
	to := c.Query("to")
	keyword := c.Query("keyword")

	emails, err := h.service.Search(
		subject,
		from,
		to,
		keyword,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Failed to get emails",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Emails retrieved successfully",
			"data":    emails,
		},
	)
}
