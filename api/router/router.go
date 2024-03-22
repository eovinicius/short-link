package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	app := gin.Default()

	AppRoutes(app)

	app.Run("localhost:8080")
}
