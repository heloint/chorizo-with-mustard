package authentication

import (
	userDAO "api/pkg/models"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type customClaim struct {
    jwt string `json:jwt`
    jwt.RegisteredClaims
}

type loginCredens struct {
    Username string `json:username`
    Password string `json:password`
}

type confirmationRes struct {
    Result bool `json:result`
}

type sessionAuth struct {
    IsLoggedIn bool `json:isLoggedIn`
    User userDAO.User `json:user`
}

func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func checkPasswordHash(password string, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func setJWTCookie(context *gin.Context, user userDAO.User) error {
    claims := customClaim {
        "jwt",
        jwt.RegisteredClaims{
            Issuer: strconv.Itoa(user.Id),
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    ss, err := token.SignedString([]byte("supersecretkey"))

    if err != nil {
        return err
    }

    cookie := http.Cookie {
        Name:"jwt",
        Value: ss,
        Path:"/",
        Expires: time.Now().Add(time.Hour * 24),
        HttpOnly: true,
    }

    http.SetCookie(context.Writer, &cookie)

    return nil
}


func DoLoginUser(c *gin.Context) {
    var newLoginCredens loginCredens
    var foundUser userDAO.User
    var err error

    // Call BindJSON to bind the received JSON to newLoginCredens.
    if err = c.BindJSON(&newLoginCredens); err != nil {
        return
    }
    foundUser, err = userDAO.GetByUsername (
        newLoginCredens.Username,
    )

    if err != nil {
        c.IndentedJSON(http.StatusUnauthorized, userDAO.User{})
        return
    }

    hashCheck := checkPasswordHash(newLoginCredens.Password, foundUser.Password)

    if hashCheck == false {
        if newLoginCredens.Password != foundUser.Password {
            c.IndentedJSON(http.StatusUnauthorized, userDAO.User{})
            return
        }
    }

    foundUser.Password = ""

    err = setJWTCookie(c, foundUser)

    if err != nil {
        c.IndentedJSON(http.StatusUnauthorized, userDAO.User{})
        return
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

    hashedPass, err := hashPassword(newUser.Password)

    if err != nil {
        log.Println(err)
        insertionRes.Result = false
    }

    newUser.Password = hashedPass
    err = userDAO.InsertNewUser(newUser)

    if err != nil {
        log.Println(err)
        insertionRes.Result = false
        c.IndentedJSON(http.StatusConflict, insertionRes)
        return
    }

    c.IndentedJSON(http.StatusCreated, insertionRes)
}

func DoUserProfile(c *gin.Context) {
    authToken := ""
    insertionRes := sessionAuth { 
        IsLoggedIn: false,
        User: userDAO.User {},
    }

    if val, ok := c.Request.Header["Authorization"]; ok {
        authToken = strings.Split(val[0], " ")[1]
    }

    claims := jwt.MapClaims{}
    token, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte("supersecretkey"), nil
    })

    if err != nil || !token.Valid {
        c.IndentedJSON(http.StatusUnauthorized, insertionRes)
        return
    }
    userIDAsStr, err := claims.GetIssuer()

    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, insertionRes)
        return
    }

    userID, err := strconv.Atoi(userIDAsStr)

    if err != nil {
        c.IndentedJSON(http.StatusInternalServerError, insertionRes)
        return
    }

    user, err := userDAO.GetByID(userID)

    insertionRes = sessionAuth { 
        IsLoggedIn: false,
        User: user,
    }

    c.IndentedJSON(http.StatusCreated, insertionRes)
}

