package app

import (
	"MyApp/internal/app/router"

	"github.com/gin-gonic/gin"
)

func InitGinEngine(r router.IRouter) *gin.Engine {
	app := gin.New()
	r.Register(app)
	return app
}
