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

type confirmationRes struct {
    Result bool `json:result`
}

func DoLoginUser(c *gin.Context) {
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

func DoRegisterUser(c *gin.Context) {
    var err error
    var newUser userDAO.User
    insertionRes := confirmationRes {Result: true}

    // Call BindJSON to bind the received JSON to newLoginCredens.
    if err := c.BindJSON(&newUser); err != nil {
        insertionRes.Result = false
    }

    err = userDAO.InsertNewUser(newUser)

    if err != nil {
        insertionRes.Result = false
    }

    c.IndentedJSON(http.StatusCreated, insertionRes)
}
