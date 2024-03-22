package handler

import (
	"net/http"

	"github.com/eovinicius/short-link/api/config"
	"github.com/eovinicius/short-link/api/schema"
	"github.com/gin-gonic/gin"
)

func CreateShortLink(c *gin.Context) {
	link := schema.NewShortLink()

	err := c.BindJSON(&link)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db := config.DbPostgres()

	db.Exec("INSERT INTO short_links (id, code, original_url, created_at) VALUES ($1, $2, $3, $4)", link.ID, link.Code, link.OriginalUrl, link.CreatedAt)

	c.JSON(http.StatusCreated, link.ID)
}
