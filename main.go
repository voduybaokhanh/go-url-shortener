package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/voduybaokhanh/go-url-shortener/config"
	"github.com/voduybaokhanh/go-url-shortener/models"
	"github.com/voduybaokhanh/go-url-shortener/routes"
)

func main() {
	config.ConnectDB()

	if err := config.DB.AutoMigrate(&models.User{}, &models.Link{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":3030")
}
