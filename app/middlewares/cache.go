package middlewares

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/dhxmo/shop-stop-go/config"
	"github.com/gin-gonic/gin"
)

var cache = config.NewRedis()
var memCache = config.NewInMemoryCache()

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// caching middleware checks the memCache map for small responses before checking Redis.
// If the response is small (less than 1 KB), it caches it in the memCache map instead
// of Redis. Responses larger than 1 KB are still cached in Redis. Responses to
// non-GET requests are removed from both the Redis cache and the memCache map.
func Cached() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cache == nil {
			log.Println("Cache is not available")
			c.Next()
			return
		}

		key := c.Request.URL.RequestURI()
		if c.Request.Method != "GET" {
			c.Next()

			if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
				cache.Remove(key)
				memCache.Remove(key)
			}
			return
		}

		// Check the in-memory cache first
		var data map[string]interface{}
		if err := memCache.Get(key, &data); err == nil {
			c.JSON(http.StatusOK, data)
			c.Abort()
			return
		}

		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		if err := cache.Get(key, &data); err == nil {
			if len(data) < 1024 { // Cache small responses in memory
				val, _ := json.Marshal(data)
				memCache.Set(key, val)
			} else { // Cache large responses in Redis
				val, _ := json.Marshal(w.body.Bytes())
				cache.Set(key, val)
			}
			c.JSON(http.StatusOK, data)
			c.Abort()
			return
		}

		c.Next()

		statusCode := w.Status()
		if statusCode == http.StatusOK {
			if w.body.Len() < 1024 { // Cache small responses in memory
				var data map[string]interface{}
				err := json.Unmarshal(w.body.Bytes(), &data)
				if err == nil {
					val, _ := json.Marshal(data)
					memCache.Set(key, val)
				}
			} else { // Cache large responses in Redis
				val, _ := json.Marshal(w.body.Bytes())
				cache.Set(key, val)
			}
		}
	}
}
