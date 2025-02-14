package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
)

func AuthorizationMiddleware(ctx *gin.Context) {
	tokenString := ctx.Request.Header.Get("Authorization")
	if tokenString == "" {
		fmt.Fprint(ctx.Copy().Writer, "Missing authorization header")
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := VerifyToken(tokenString)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func CreateToken(email string) (string, error) {
	expirationTime, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_TIME"))
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * time.Duration(expirationTime)).Unix(),
		},
	)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func CheckPasswordHash(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
