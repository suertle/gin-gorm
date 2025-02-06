package route

import (
	"gin-gorm/src/handler"

	"github.com/gin-gonic/gin"
)

func setupApiRoutes(api *gin.RouterGroup) {
	// playground for test some code
	api.GET("/test", handler.Test)
}
