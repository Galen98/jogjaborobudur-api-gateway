package main

import(
	 "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main(){
	r := gin.Default()

	// Setup CORS config
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://jogjaborobudur.com", "https://www.jogjaborobudur.com", "https://api.jogjaborobudur.com", "https://www.api.jogjaborobudur.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//get contoh
	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message":"teses",
		})
	})

	api := r.Group("/api/v1")
    {

	}
	r.Run(":8080")
}