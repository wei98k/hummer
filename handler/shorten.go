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
}

func Shorten(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.OriginalURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	shortCode := model.GenerateCode() // 随机生成短码

	_, err := storage.DB.Exec(`
        INSERT INTO short_links (short_code, original_url, expire_at)
        VALUES (?, ?, ?)`,
		shortCode, req.OriginalURL, req.ExpireAt,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"short_url": "http://localhost:8080/go/" + shortCode,
	})
}
