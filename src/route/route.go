package route

import (
	"fmt"
	"gin-gorm/src/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() (router *gin.Engine) {
	router = gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	// router.Use(requestlog.RequestLogHandler(db.Eddu))
	router.Use(middleware.Sentry()...)

	router.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.GET("/panic", func(c *gin.Context) {
		panic(fmt.Errorf("this is a test panic"))
	})

	// for health check
	router.GET("/health", func(c *gin.Context) { c.String(200, "OK") })

	setupApiRoutes(router.Group(`/api`))

	return
}
