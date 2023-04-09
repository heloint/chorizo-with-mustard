package jwtAuth

import (
	userDAO "api/pkg/models"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type customClaim struct {
    jwt string `json:jwt`
    jwt.RegisteredClaims
}

func SetJWTCookie(context *gin.Context, user userDAO.User) error {
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

func ExtractJWTToken(context *gin.Context) (string, error) {
    authToken := ""

    // Get auth token
    authHeader := context.Request.Header.Get("Authorization")

    if authHeader == "" {
        return authToken, errors.New("\"Authorization\" in headers is empty!")
    }
    authToken = strings.Split(authHeader, " ")[1]
    return authToken, nil
}
