package handler

import (
	"net/http"

	"github.com/eovinicius/short-link/api/config"
	"github.com/eovinicius/short-link/api/schema"
	"github.com/gin-gonic/gin"
)

func GetLinks(c *gin.Context) {
	db := config.DbPostgres()

	rows, err := db.Query("SELECT * FROM short_links")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	defer rows.Close()

	var links []*schema.ShortLink

	for rows.Next() {
		var link schema.ShortLink
		err := rows.Scan(&link.ID, &link.Code, &link.OriginalUrl, &link.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		links = append(links, &link)
	}

	c.JSON(http.StatusOK, links)
}
