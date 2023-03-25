package apiLogs

import (
    "os"
    "io"
    "log"
    "github.com/gin-gonic/gin"
)

func InitAPILogs() {
    var err error
    var logFile *os.File
    var errorLogFile *os.File
    
    if err := os.Mkdir("./logs", 0755); err != nil {
        log.Println(err)
    }

    logFile, err = os.Create("./logs/request.log")

    if err != nil {
        panic(err)
    }

    gin.DefaultWriter = io.MultiWriter(logFile)

    errorLogFile, err = os.Create("./logs/error.log")

    if err != nil {
        panic(err)
    }

    gin.DefaultErrorWriter = io.MultiWriter(errorLogFile)
}
