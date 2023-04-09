package jwtAuth

import (
    "os"
	userDAO "api/pkg/models"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = os.Getenv("JWT_KEY")

type customClaim struct {
	jwt string `json:jwt`
	jwt.RegisteredClaims
}

// Sets the JWT token as a httpOnly cookie for the client.
func SetJWTCookie(context *gin.Context, user userDAO.User) error {
	claims := customClaim{
		"jwt",
		jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(user.Id),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		return err
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    ss,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	http.SetCookie(context.Writer, &cookie)
	return nil
}

// Extracts the JWT token from the headers.
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

// Returns the claims from a token if it's valid, else an error.
func GetTokenClaims(authToken string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil || !token.Valid {
		return jwt.MapClaims{}, err
	} else if !token.Valid {
		return jwt.MapClaims{}, errors.New("The given token is invalid!")
	}

	return claims, nil
}

// Expires the jwt cookie on the client.
func ExpireJWTCookie(context *gin.Context) {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(context.Writer, &cookie)
}
