package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKey         = []byte(os.Getenv("JWT_SECRET_KEY"))
	expirationTime, _ = strconv.Atoi(os.Getenv("JWT_EXPIRATION_TIME"))
)

func LoginHandler(c *gin.Context) {
	c.Request.Header.Set("Content-Type", "application/json")

	//var u handler.User
	//json.NewDecoder(c.Request.Response.Body).Decode(&u)
	//fmt.Printf("The user request value %v", u)

	user := c.PostForm("username")
	pass := c.PostForm("password")

	if user != "user" || pass != "password" {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(c.Copy().Writer, "Invalid credentials")
	}

	tokenString, err := createToken(user)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		fmt.Errorf("No username found")
	}
	c.Writer.WriteHeader(http.StatusOK)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func AuthorizationMiddleware(c *gin.Context) {
	c.Request.Header.Set("Content-Type", "application/json")
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		fmt.Fprint(c.Copy().Writer, "Missing authorization header")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := verifyToken(tokenString)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * time.Duration(expirationTime)).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
