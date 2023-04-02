package main

import (
    _"api/pkg/apiLogs"
    _"api/pkg/models"
    "api/pkg/controllers"
    "github.com/gin-gonic/gin"
    _"github.com/gin-contrib/cors"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, access-control-allow-origin, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
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
    /* config := cors.DefaultConfig()
    // config.AllowAllOrigins = true
    config.AllowOrigins = []string{"http://localhost:3000"}
    config.AllowCredentials = true
    config.AllowHeaders = []string{"Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers"}
    config.AllowOriginFunc = func(origin string) bool {
        return origin == "http://localhost:3000"
    }
    r.Use(cors.New(config)) */
    r.Use(CORSMiddleware())
    r.POST("/login", authentication.DoLoginUser)
    r.POST("/register", authentication.DoRegisterUser)
    r.Run("0.0.0.0:8000")
}
