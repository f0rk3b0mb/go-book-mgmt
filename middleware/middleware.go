package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"

	initializer "github.com/f0rk3b0mb/go-book-mgmt/initializers"
	model "github.com/f0rk3b0mb/go-book-mgmt/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Validate(c *gin.Context) {
	authtoken, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "please login"})
		c.AbortWithStatus(401)
		return
	}

	token, err := jwt.Parse(authtoken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		var user model.Users

		initializer.DB.First(&user, claims["id"])
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "please login"})
			c.AbortWithStatus(401)
			return
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "please login"})
		c.AbortWithStatus(401)
		return
	}

}
