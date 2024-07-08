// middleware/logging.go

package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware logs details of each request and its execution time
func LoggingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Start timer
        start := time.Now()

        // Process request
        c.Next()

        // Calculate latency
        latency := time.Since(start)

        // Log request details
        log.Printf("[%s] %s %s %s %v\n",
            time.Now().Format("2006-01-02 15:04:05"),
            c.Request.Method,
            c.Request.URL.Path,
            c.Request.RemoteAddr,
            latency,
        )
    }
}
