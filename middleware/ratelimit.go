package middleware

import (
	"sync"
	"time"

	"2_TaskManager/config"

	"github.com/gin-gonic/gin"
)

type visitor struct {
	Count     int
	LastReset time.Time
}

var visitors = make(map[string]*visitor)
var mu sync.Mutex

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()

		v, exists := visitors[ip]
		if !exists {
			visitors[ip] = &visitor{
				Count:     1,
				LastReset: time.Now(),
			}
			mu.Unlock()
			c.Next()
			return
		}

		// Reset counter every 10 seconds
		if time.Since(v.LastReset) > config.RateLimitWindow {
			v.Count = 1
			v.LastReset = time.Now()
			mu.Unlock()
			c.Next()
			return
		}

		if v.Count >= config.RateLimitRequests {
			mu.Unlock()
			c.JSON(429, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}

		v.Count++
		mu.Unlock()

		c.Next()
	}
}
