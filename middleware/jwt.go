package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret string

func maskSecret(secret string) string {
	if len(secret) <= 10 {
		return secret // Jika kurang dari 10 karakter, tampilkan apa adanya
	}
	return secret[:5] + "..." + secret[len(secret)-5:]
}

// init() akan dijalankan saat package middleware pertama kali di-load
func init() {
	jwtSecret = os.Getenv("JWT_SECRET") // Ambil dari environment variable
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is required") // Jika kosong, langsung fatal error
	} else {
		log.Println("Loaded JWT_SECRET:", maskSecret(jwtSecret)) // Log untuk debugging
	}
}


// Middleware JWT dengan secret yang sudah otomatis dimuat
func EjourneyMiddlewareUserJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Cek algoritma harus HS512
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Pastikan signature diverifikasi dengan secret
			return []byte(jwtSecret), nil
		}, jwt.WithValidMethods([]string{"HS512"}))

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		// Ambil claims dari token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token claims"})
			c.Abort()
			return
		}

		// Cek expiry token
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Token has expired"})
				c.Abort()
				return
			}
		}

		// Simpan claims ke context
		c.Set("userID", claims["sub"])
		c.Set("claims", claims)
		c.Next()
	}
}