package routes

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/voduybaokhanh/go-url-shortener/controllers"
)

func SetupRoutes(r *gin.Engine) {
	// Public
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/r/:code", controllers.Redirect)

	// Protected
	protected := r.Group("/")
	protected.Use(AuthMiddleware())
	{
		protected.POST("/shorten", controllers.CreateLink)
		protected.GET("/links", controllers.GetLinks)
	}
}

// Middleware để check JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		secret := os.Getenv("JWT_SECRET")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["user_id"].(float64))
		c.Set("user_id", userID)

		c.Next()
	}
}
