package main

import (
    _"api/pkg/apiLogs"
    _"api/pkg/models"
    "api/pkg/controllers"
    "github.com/gin-gonic/gin"
)

func main() {

    // Start writing the SDTOUT & STDERR to log files.
    // apiLogs.InitAPILogs()

    r := gin.Default()
    r.POST("/login", authentication.DoLoginUser)
    r.POST("/register", authentication.DoRegisterUser)
    r.Run("0.0.0.0:8000")

}
