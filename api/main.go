package main

import (
	_ "api/pkg/apiLogs"
	"api/pkg/controllers"
	_ "api/pkg/models"
	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Set API headers
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header(
            "Access-Control-Allow-Headers", 
            `Content-Type
             access-control-allow-origin
             Content-Length
             Accept-Encoding
             X-CSRF-Token
             Authorization
             accept
             origin
             Cache-Control
             X-Requested-With`,
         )
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {

	// Start writing the SDTOUT & STDERR to log files.
	// apiLogs.InitAPILogs()


	r := gin.Default()

    // Apply defined headers.
	r.Use(CORSMiddleware())

    // Routes
	r.POST("/login", authentication.DoLoginUser)
	r.POST("/register", authentication.DoRegisterUser)
	r.GET("/profile", authentication.DoUserProfile)

	r.Run("0.0.0.0:8000")
}
