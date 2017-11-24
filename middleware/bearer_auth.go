package middleware

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//BearerAuthMiddleware to add JWT security to gin-gonic resources
func BearerAuthMiddleware(ptr gin.HandlerFunc, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		//setCORSEnabled(c)

		tokenString := strings.SplitN(c.GetHeader("Authorization"), " ", 2)
		if len(tokenString) != 2 || tokenString[0] != "Bearer" {
			c.String(http.StatusUnauthorized, "Bearer header required")
			return
		}

		token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validateUser the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return publicKey, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			expires, ok := claims["exp"]
			if !ok {
				sendUnauthorized(c, errors.New("Token not valid"))
				return
			}

			if getTokenRemainingValidity(expires) == -1 {
				sendUnauthorized(c, errors.New("Token validity expired"))
				return
			}

			c.Set("claims", claims)
		} else {
			sendUnauthorized(c, err)
			return
		}
		ptr(c)
	}
}

func sendUnauthorized(c *gin.Context, err error) {
	fmt.Println(err.Error())
	c.String(http.StatusUnauthorized, "Token not valid: ", err.Error())
	return
}

func getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds())
		}
	}
	return -1
}
