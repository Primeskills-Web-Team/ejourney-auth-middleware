package middleware

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
)

var jwtSecret string

// init() akan dijalankan saat package middleware pertama kali di-load
func init() {
	jwtSecret = os.Getenv("JWT_SECRET") // Ambil dari environment variable
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is required") // Jika kosong, langsung fatal error
	} else {
		log.Println("Loaded JWT_SECRET:", jwtSecret) // Log untuk debugging
	}
}

// Middleware JWT dengan secret yang sudah otomatis dimuat
func EjourneyMiddlewareUserJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Gunakan jwtSecret untuk validasi token
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Validasi token (contoh sederhana)
		if token != "Bearer "+jwtSecret {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
