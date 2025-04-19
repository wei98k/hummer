package handler

import (
	"github.com/gin-gonic/gin"
	"hummer/model"
	"hummer/storage"
	"net/http"
	"time"
)

type ShortenRequest struct {
	OriginalURL string    `json:"original_url"`
	ExpireAt    time.Time `json:"expire_at,omitempty"`
	Title       string    `json:"title,omitempty"`
}

func Shorten(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.OriginalURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var shortCode string
	for {
		shortCode = model.GenerateCode()
		var exists int
		err := storage.DB.QueryRow(`SELECT COUNT(1) FROM short_links WHERE short_code = ?`, shortCode).Scan(&exists)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
			return
		}
		if exists == 0 {
			break
		}
	}

	_, err := storage.DB.Exec(`
        INSERT INTO short_links (short_code, original_url, expire_at, title)
        VALUES (?, ?, ?, ?)`,
		shortCode, req.OriginalURL, req.ExpireAt, req.Title,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"short_url": "http://go.123191.xyz/" + shortCode,
		"title":     req.Title,
	})
}
