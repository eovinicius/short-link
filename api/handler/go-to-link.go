package handler

import (
	"context"
	"net/http"

	"github.com/eovinicius/short-link/api/config"
	"github.com/eovinicius/short-link/api/schema"
	"github.com/gin-gonic/gin"
)

func GoToLink(c *gin.Context) {
	code := c.Param("code")

	db := config.DbPostgres()

	row := db.QueryRow("SELECT * FROM short_links WHERE code = $1", code)

	var link schema.ShortLink

	err := row.Scan(&link.ID, &link.Code, &link.OriginalUrl, &link.CreatedAt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "erro ao buscar link"})
		return
	}

	redis := config.DbRedis()

	ctx := context.Background()

	redis.ZIncrBy(ctx, "hits", 1, link.Code)

	c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
}
