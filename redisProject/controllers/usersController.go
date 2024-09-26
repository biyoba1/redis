package controllers

import (
	"github.com/biyoba1/redisProject/initializer"
	"github.com/biyoba1/redisProject/internal/models"
	"github.com/biyoba1/redisProject/redis"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

type Authorization interface {
	SignUp(c *gin.Context)
	Login(c *gin.Context)
	Validate(c *gin.Context)
	//Logout(c *gin.Context)
}

type Auth struct{}

func (a *Auth) SignUp(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	user := models.Person{Email: body.Email, Password: string(hash)}

	result := initializer.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	err = redis.CacheUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to cache user",
		})
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (a *Auth) Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	var person models.Person
	initializer.DB.First(&person, "email = ?", body.Email)

	if person.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(person.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": person.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "couldn't get token",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{})
}

func (a *Auth) Validate(c *gin.Context) {
	user, _ := c.Get("person")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
