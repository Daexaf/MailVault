package route

import (
	"github.com/daexaf/mailvault/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterBranchRoutes(router *gin.RouterGroup, handler *handler.BranchHandler) {
	router.POST("/", handler.Create)
}
