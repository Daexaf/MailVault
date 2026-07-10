package route

import (
	"github.com/daexaf/mailvault/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterMailAccountRoutes(router *gin.RouterGroup, handler *handler.MailAccountHandler) {
	router.POST("/", handler.Create)
	router.POST("/:id/test", handler.TestConnection)
}
