package route

import (
	"github.com/daexaf/mailvault/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterEmailRoutes(
	router *gin.RouterGroup,
	handler *handler.EmailHandler,
) {
	router.GET("", handler.Search)
}
