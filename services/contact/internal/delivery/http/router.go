package http

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (d *Delivery) initRouter() *gin.Engine {
	if viper.GetBool("IS_PRODUCTION") {
		switch strings.ToUpper(strings.TrimSpace(viper.GetString("LOG_LEVEL"))) {
		case "DEBUG":
			gin.SetMode(gin.DebugMode)
		default:
			gin.SetMode(gin.ReleaseMode)
		}
	} else {
		gin.SetMode(gin.DebugMode)
	}

	var router = gin.New()
	router.Use(checkAuth)

	return router
}
