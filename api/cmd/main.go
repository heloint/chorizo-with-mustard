package main

import (
    "fmt"
    "net/http"
    _"api/pkg/apiLogs"
    "api/pkg/models"
    "api/pkg/controllers"
    "github.com/gin-gonic/gin"
)

func main() {

    // Start writing the SDTOUT & STDERR to log files.
    // apiLogs.InitAPILogs()

    users := userDAO.GetAll()

    fmt.Println(users)

    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
      c.JSON(http.StatusOK, users)
    })
    r.POST("/login", authentication.DoLogin)
    r.Run("0.0.0.0:8000")

}
