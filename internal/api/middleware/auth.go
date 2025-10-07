package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	error_message "github.com/jonathanmoreiraa/2cents/internal/domain/error"
	"github.com/jonathanmoreiraa/2cents/pkg/util"

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
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Token inválido! Adicione o token no cabeçalho",
		})
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := VerifyToken(tokenString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Token inválido! Realize o login novamente",
		})
		return
	}
}

func CreateToken(id int) (string, error) {
	expirationTime, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_TIME"))
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  id,
			"exp": time.Now().Add(time.Hour * time.Duration(expirationTime)).Unix(),
		},
	)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", util.ErrorWithMessage(err, error_message.ErrCreateToken)
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, util.ErrorWithMessage(nil, fmt.Sprintf("%s: %v", error_message.ErrMethodLogin, token.Header["alg"]))
		}
		return secretKey, nil
	})
	if err != nil {
		return util.ErrorWithMessage(err, error_message.ErrInvalidToken)
	}

	if !token.Valid {
		return util.ErrorWithMessage(nil, error_message.ErrInvalidToken)
	}

	return nil
}

func CheckPasswordHash(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ExtractUserId(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, util.ErrorWithMessage(err, error_message.ErrDecodeToken)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, util.ErrorWithMessage(nil, error_message.ErrDecodeToken)
	}

	id, ok := claims["id"].(float64)
	if !ok {
		return 0, util.ErrorWithMessage(nil, error_message.ErrDecodeToken)
	}

	return int(id), nil
}
