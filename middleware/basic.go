package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ValidateUserCallback is used to validate a user sent by http basic authentication header
type ValidateUserCallback func(c *gin.Context, user string, passwd string) bool

// AuthMiddleware enables Basic authentication security to a gin.HandlerFunc
func AuthMiddleware(ptr gin.HandlerFunc, validateUser ValidateUserCallback) gin.HandlerFunc {
	return func(c *gin.Context) {
		//setCORSEnabled(c)

		auth := strings.SplitN(c.GetHeader("Authorization"), " ", 2)
		if len(auth) != 2 || auth[0] != "Basic" {
			c.Status(http.StatusUnauthorized)
			return
		}
		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !validateUser(c, pair[0], pair[1]) {
			c.Status(http.StatusUnauthorized)
			return
		}
		ptr(c)
	}
}
