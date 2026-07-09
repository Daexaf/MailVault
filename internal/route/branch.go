package route

import (
	"github.com/daexaf/mailvault/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterBranchRoutes(router *gin.RouterGroup, handler *handler.BranchHandler) {
	router.POST("/", handler.Create)
	router.GET("/", handler.FindAll)
	router.GET("/:id", handler.FindByID)
}

// func FindAllBranchesRoutes(router *gin.RouterGroup, handler *handler.BranchHandler) {
// 	router.GET("/", handler.FindAll)
// }
