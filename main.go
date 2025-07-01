package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://jogjaborobudur.com", "https://www.jogjaborobudur.com", "https://api.jogjaborobudur.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Group API v1
	api := r.Group("/api/v1")
	{
		api.Any("/data-booking", proxyHandler("https://jogjaborobudur.com/api/v1/data-booking"))
	}

	// Jalankan di port 8080
	r.Run(":8080")
}

// Proxy handler
func proxyHandler(targetURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Buat request baru ke target backend
		req, err := http.NewRequest(c.Request.Method, targetURL, c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}

		// Copy headers dari request asli
		for key, values := range c.Request.Header {
			for _, v := range values {
				req.Header.Add(key, v)
			}
		}

		// Kirim request ke backend
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to reach backend"})
			return
		}
		defer resp.Body.Close()

		// Copy response dari backend ke client
		for key, values := range resp.Header {
			for _, v := range values {
				c.Writer.Header().Add(key, v)
			}
		}
		c.Status(resp.StatusCode)
		io.Copy(c.Writer, resp.Body)
	}
}
