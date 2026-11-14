package middleware

import (
	"net/http"
	"time"

	"asri-backend/internal/store"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// limitKey prefers the login identity: shared school IPs make IP-only
// limits class-wide. Pre-auth callers fall back to IP.
func limitKey(c *gin.Context) string {
	base := c.FullPath()
	if v, ok := c.Get("claims"); ok {
		if m, ok := v.(jwt.MapClaims); ok {
			if n, _ := m["name"].(string); n != "" {
				if code, _ := m["code"].(string); code != "" {
					return base + "|n:" + code + ":" + n
				}
				return base + "|n:" + n
			}
			if s, _ := m["sub"].(string); s != "" {
				return base + "|u:" + s
			}
		}
	}
	return base + "|ip:" + c.ClientIP()
}

func check(r *store.Redis, c *gin.Context, limit int, window time.Duration, failClosed bool) {
	ok, err := r.Allow(c.Request.Context(), limitKey(c), limit, window)
	if err != nil || !ok {
		if err != nil && !failClosed {
			c.Next() // fail open — throttling only; callers stay authoritative
			return
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"error": "try again shortly"})
			return
		}
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded, try again shortly"})
		return
	}
	c.Next()
}

// RateLimit throttles per caller; survives Redis outages open.
func RateLimit(r *store.Redis, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) { check(r, c, limit, window, false) }
}

// RateLimitStrict throttles per caller and refuses when Redis is down.
// Use for the PIN gate: the code is the only credential, so unthrottled
// guessing during an outage is worse than a short outage.
// shortcut: Redis-down blocks new PIN entries; DB-backed validate itself
// keeps working once Redis returns. A local in-memory window removes even that.
func RateLimitStrict(r *store.Redis, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) { check(r, c, limit, window, true) }
}
