package middleware

import "github.com/gin-gonic/gin"

// CORSManual is fallback permissive CORS for dev (like Express app.use(cors())).
// gin-contrib/cors already handles preflight, this just ensures headers.
func CORSManual() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
