package middleware

import (
	"net/http"
	"strings"
	"edutest/pkg/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Next()
	}
}

func AuthMiddleware(jwtKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, model.Error{
				Message: "Authorization header is missing",
				Error:   "unauthorized",
			})
			c.Abort()
			return
		}

		var tokenStr string

		// Agar header Bearer bilan boshlangan bo‘lsa
		if strings.HasPrefix(strings.ToLower(authHeader), "bearer ") {
			tokenStr = strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
		} else {
			// Aks holda Swagger holati: faqat token yuborilgan
			tokenStr = strings.TrimSpace(authHeader)
		}

		// JWTni tekshirish
		claims := &model.Claim{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, model.Error{
				Message: "Invalid or expired token",
				Error:   err.Error(),
			})
			c.Abort()
			return
		}

		// Foydalanuvchi ID sini contextga saqlash
		c.Set("user_id", claims.Items["user_id"])
		c.Next()
	}
}
