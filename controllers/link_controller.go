package controllers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/voduybaokhanh/go-url-shortener/config"
	"github.com/voduybaokhanh/go-url-shortener/models"
)

// Helper để lấy user_id từ context (JWT middleware sẽ set)
func getUserID(c *gin.Context) uint {
	uid, _ := c.Get("user_id")
	return uid.(uint)
}

// POST /shorten
func CreateLink(c *gin.Context) {
	var body struct {
		URL string `json:"url" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortCode := RandString(6)
	link := models.Link{OriginalURL: body.URL, ShortCode: shortCode, UserID: getUserID(c)}

	if err := config.DB.Create(&link).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not create link"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": "/r/" + shortCode})
}

// GET /r/:code → redirect
func Redirect(c *gin.Context) {
	code := c.Param("code")
	var link models.Link
	if err := config.DB.Where("short_code = ?", code).First(&link).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}
	c.Redirect(http.StatusFound, link.OriginalURL)
}

// GET /links → danh sách link của user
func GetLinks(c *gin.Context) {
	var links []models.Link
	config.DB.Where("user_id = ?", getUserID(c)).Find(&links)
	c.JSON(http.StatusOK, links)
}

// Helper: random string
func RandString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
