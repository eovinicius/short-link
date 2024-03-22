package router

import (
	"github.com/eovinicius/short-link/api/handler"
	"github.com/gin-gonic/gin"
)

func AppRoutes(routes *gin.Engine) {

	routes.POST("/api/v1/links", handler.CreateShortLink)
	routes.GET("/api/v1/links", handler.GetLinks)
	routes.GET("/:code", handler.GoToLink)
}
