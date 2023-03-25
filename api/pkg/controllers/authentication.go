package authentication

import (
    "net/http"
    "api/pkg/models"
    "github.com/gin-gonic/gin"
)

type loginCredens struct {
    Username string `json:username`
    Password string `json:password`
}

func DoLogin(c *gin.Context) {
    var newLoginCredens loginCredens

    // Call BindJSON to bind the received JSON to newLoginCredens.
    if err := c.BindJSON(&newLoginCredens); err != nil {
        return
    }
    var foundUser userDAO.User = userDAO.GetByUsernameAndPassword (
        newLoginCredens.Username,
        newLoginCredens.Password,
    )

    c.IndentedJSON(http.StatusCreated, foundUser)
}
