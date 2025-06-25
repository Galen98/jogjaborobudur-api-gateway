package main

import(
	 "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main(){
	r := gin.Default()

	// Setup CORS config
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3000", "https://namadomainmu.com"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	//get contoh
	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message":"tes",
		})
	})

	api := r.Group("/api/v1")
    {

	}
	r.Run(":8080")
}