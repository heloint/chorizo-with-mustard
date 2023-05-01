package main

import (
	_ "api/pkg/apiLogs"
	"api/pkg/controllers/authentication"
	"api/pkg/controllers/downloads"
	_ "api/pkg/models"
	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
    _"log"
)

// Set API headers
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header(
            "Access-Control-Allow-Headers", 
            `Content-Type,
             access-control-allow-origin,
             Content-Length,
             Accept-Encoding,
             X-CSRF-Token,
             Authorization,
             accept,
             origin,
             Cache-Control,
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

	router := gin.Default()

    // Apply defined headers.
	router.Use(CORSMiddleware())

	router.POST("/login", authentication.DoLoginUser)
	router.POST("/register", authentication.DoRegisterUser)
	router.GET("/profile", authentication.DoUserProfile)
	router.GET("/logout", authentication.DoLogout)

	router.GET("/downloads", downloads.ListAllFiles)
    router.GET("/downloads/:file", downloads.DownloadFile)

    /* r. GET("/user/:username", func (c *gin.Context) {
        username := c.Param("username")
        log.Println(username)
    }) */

	router.Run("0.0.0.0:8000")
}
