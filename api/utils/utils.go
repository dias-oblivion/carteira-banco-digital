package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dias-oblivion/carteira-banco-digital/api/config"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func CreateJWTToken(userId int64, userEmail, userRole string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["userEmail"] = userEmail
	claims["role"] = userRole
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.API_SECRET))
}

func VerifyJWTToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.API_SECRET), nil
	})
	if err != nil {
		return err
	}

	if parsedToken.Valid {
		return nil
	}

	return err
}

func GetCurrentUserInfo(ctx *gin.Context) (UserInfo, error) {

	bearerToken := strings.Split(ctx.Request.Header.Get("Authorization"), " ")[1]

	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return UserInfo{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, err := strconv.ParseInt(fmt.Sprintf("%v", claims["userId"]), 10, 64)
		if err != nil {
			return UserInfo{}, err
		}
		return UserInfo{
			ID:    userId,
			Email: string(fmt.Sprintf("%v", claims["userEmail"])),
			Role:  string(fmt.Sprintf("%v", claims["userRole"])),
		}, nil
	}
	return UserInfo{}, nil
}

func GetResponseJson(response *http.Response, target interface{}) error {
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(target)
}
