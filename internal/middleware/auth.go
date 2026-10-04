package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


const (
	KindTeacher = "teacher"
	KindSession = "session"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header must be Bearer token"})
			return
		}
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "dev-secret-change-me"
		}

		token, err := jwt.Parse(parts[1], func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			return
		}

		kind, _ := claims["kind"].(string)
		if kind != KindTeacher && kind != KindSession {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unknown token kind"})
			return
		}
		c.Set("tokenKind", kind)
		if sub, ok := claims["sub"].(string); ok {
			c.Set("userID", sub)
		}
		if chapterID, ok := claims["chapter_id"].(string); ok {
			c.Set("chapterID", chapterID)
		}
		c.Set("claims", claims)

		c.Next()
	}
}


func RequireKind(kind string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if k, _ := c.Get("tokenKind"); k != kind {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
			return
		}
		c.Next()
	}
}

func SessionChapter(c *gin.Context) string {
	if k, _ := c.Get("tokenKind"); k != KindSession {
		return ""
	}
	chapterID, _ := c.Get("chapterID")
	id, _ := chapterID.(string)
	return id
}
