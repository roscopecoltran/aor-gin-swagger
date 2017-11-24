package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
	Refs:
	- https://github.com/bradberger/gologin/blob/master/handlers.go
	- https://github.com/xenu256/PiggyBob/blob/master/main.go
*/

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func SetTokenHandler(c *gin.Context) {
	expireToken := time.Now().Add(time.Hour * 24).Unix()
	expireCookie := time.Now().Add(time.Hour * 24)

	claims := MyCustomClaims{
		"moxiaobai",
		jwt.StandardClaims{
			ExpiresAt: expireToken,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, _ := token.SignedString([]byte("secret"))

	cookie := http.Cookie{Name: "Auth", Value: signedToken, Expires: expireCookie, HttpOnly: true}
	http.SetCookie(c.Writer, &cookie)
}
