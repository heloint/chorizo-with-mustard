package main

import (
    "fmt"
    "api/pkg/models"
)
 
func main() {
    users := user.GetAllUsers()

    fmt.Println(users)

}
