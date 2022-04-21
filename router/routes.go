package router

import (
	"go_cli/demo/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLog(), logger.GinRecover(true))
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ready")
	})
	return r
}
