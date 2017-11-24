package middleware

/*
import (
	"io/ioutil"
	"log"
	"time"

	"github.com/roscopecoltran/admin-on-rest-server/server/config"
	"github.com/roscopecoltran/admin-on-rest-server/server/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var cfg config.ServiceConfig

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(gorm.DB)
		cfg = config.Config
		// config = c.MustGet("config").(config.ServiceConfig)
		tokenString := c.Request.Header.Get("X-Authentication-Token")

		if tokenString == "" {
			response := make(map[string]string)
			response["error"] = "Invalid credentials."
			c.JSON(401, response)
			c.Abort()
		} else {
			token, err := jwt.Parse(tokenString, verify)

			if err != nil {
				log.Fatal(err)
			} else {
				if token.Valid {

					now := time.Now().UTC()
					expiration, err := time.Parse(time.RFC3339, token.Claims["life"].(string))

					if err != nil {
						log.Fatal(err)
					}

					if now.After(expiration) {
						response := make(map[string]string)
						response["error"] = "Invalid credentials."
						c.JSON(401, response)
						c.Abort()
					}

					var user models.User

					db.First(&user, token.Claims["id"])

					if user.Username == "" {
						response := make(map[string]string)
						response["error"] = "Someone fucked up."
						c.JSON(500, response)
						c.Abort()
					} else {
						c.Set("consumer", user)
						c.Next()
					}
				} else {
					response := make(map[string]string)
					response["error"] = "Someone fucked up."
					c.JSON(500, response)
					c.Abort()
				}
			}
		}
	}
}

func verify(*jwt.Token) (interface{}, error) {
	key, err := ioutil.ReadFile(config.Config.JWT.PublicKey)
	if err != nil {
		log.Fatal(err)
	}
	return key, nil
}
*/
