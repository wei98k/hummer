package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"hummer/storage"
	"net/http"
	"time"
)

func Redirect(c *gin.Context) {
	code := c.Param("short_code")
	var url string
	var expireAt sql.NullTime

	err := storage.DB.QueryRow(`
        SELECT original_url, expire_at FROM short_links WHERE short_code = ?`, code).
		Scan(&url, &expireAt)

	if err == sql.ErrNoRows {
		c.String(http.StatusNotFound, "Link not found")
		return
	} else if err != nil {
		c.String(http.StatusInternalServerError, "DB error")
		return
	}

	if expireAt.Valid && expireAt.Time.Before(time.Now()) {
		c.String(http.StatusGone, "Link expired")
		return
	}

	// 记录点击日志
	userAgent := c.Request.UserAgent()
	referer := c.Request.Referer()
	ip := c.ClientIP()
	acceptLang := c.Request.Header.Get("Accept-Language")

	go func() {
		_, _ = storage.DB.Exec(`
			INSERT INTO click_logs (short_code, user_agent, referer, ip_address, accept_lang)
			VALUES (?, ?, ?, ?, ?)`,
			code, userAgent, referer, ip, acceptLang)
	}()

	storage.DB.Exec(`UPDATE short_links SET click_count = click_count + 1 WHERE short_code = ?`, code)
	c.Redirect(http.StatusFound, url)
}
