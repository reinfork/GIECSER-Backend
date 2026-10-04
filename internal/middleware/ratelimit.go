package middleware

import (
	"net/http"
	"time"

	"asri-backend/internal/store"

	"github.com/gin-gonic/gin"
)


func RateLimit(r *store.Redis, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ok, err := r.Allow(c.Request.Context(), c.ClientIP()+":"+c.FullPath(), limit, window)
		if err != nil || !ok {
			if err != nil {
				c.Next() // ponytail: fail open — see above
				return
			}
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded, try again shortly"})
			return
		}
		c.Next()
	}
}
