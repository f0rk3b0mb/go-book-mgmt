package routes

import (
	"fmt"
	"net/http"
	"os"

	initializer "github.com/f0rk3b0mb/go-book-mgmt/initializers"
	model "github.com/f0rk3b0mb/go-book-mgmt/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func View(c *gin.Context) {
	var book []model.Book
	initializer.DB.Find(&book)
	c.JSON(http.StatusOK, gin.H{"books": book})

}

func Addbook(c *gin.Context) {
	var book model.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid json"})
		return
	}

	newbook := model.Book{Name: book.Name, Author: book.Author}

	result := initializer.DB.Create(&newbook)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "book added succesfully"})

}

func Searchbook(c *gin.Context) {
	id := c.Param("id")
	var book model.Book

	initializer.DB.First(&book, id)

	if book.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "book not found"})

	} else {
		c.JSON(http.StatusOK, gin.H{"success": book})
	}

}

func Delete(c *gin.Context) {

	id := c.Param("id")

	var book model.Book

	initializer.DB.Delete(&book, id)

	if book.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "book not found"})

	} else {
		c.JSON(http.StatusOK, gin.H{"success": "book deleted"})
	}

}

func Login(c *gin.Context) {
	var user model.Users

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid json"})
		return

	}

	var check model.Users

	initializer.DB.First(&check, "email = ?", user.Email)
	if check.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(check.Password), []byte(user.Password))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  "jwt",
		"id":   check.ID,
		"user": check.Username,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.String(http.StatusInternalServerError, "failed to generate token")
		return
	}

	fmt.Println(err)

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*12*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"success": "login successful"})

}

func Register(c *gin.Context) {
	var user model.Users

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid json"})
		return

	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	fmt.Println(user.Username, user.Email, user.Password)

	newuser := model.Users{Username: user.Username, Password: string(hash), Email: user.Email}

	result := initializer.DB.Create(&newuser)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "user registered"})

}
