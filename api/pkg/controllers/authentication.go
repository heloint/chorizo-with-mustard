package authentication

import (
    "log"
    "net/http"
    "api/pkg/models"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

type loginCredens struct {
    Username string `json:username`
    Password string `json:password`
}

type confirmationRes struct {
    Result bool `json:result`
}

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func checkPasswordHash(password string, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func DoLoginUser(c *gin.Context) {
    var newLoginCredens loginCredens

    // Call BindJSON to bind the received JSON to newLoginCredens.
    if err := c.BindJSON(&newLoginCredens); err != nil {
        return
    }

    var foundUser userDAO.User = userDAO.GetByUsername (
        newLoginCredens.Username,
    )

    hashCheck := checkPasswordHash(newLoginCredens.Password, foundUser.Password)
    if hashCheck == false {
        if newLoginCredens.Password != foundUser.Password {
            c.IndentedJSON(http.StatusUnauthorized, userDAO.User{})
            return
        }
    }

    c.IndentedJSON(http.StatusCreated, foundUser)
}

func DoRegisterUser(c *gin.Context) {
    var err error
    var newUser userDAO.User
    insertionRes := confirmationRes {Result: true}

    // Call BindJSON to bind the received JSON to newLoginCredens.
    if err := c.BindJSON(&newUser); err != nil {
        log.Println(err)
        insertionRes.Result = false
    }

    err = userDAO.InsertNewUser(newUser)

    if err != nil {
        log.Println(err)
        insertionRes.Result = false
    }

    c.IndentedJSON(http.StatusCreated, insertionRes)
}
