package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	log.Logger = log.With().Caller().Logger()
}

var secretKey = []byte("secret-key")

func CreateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func VerifyToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func GetStrValFromContext(c *gin.Context, value string) (string, bool) {
	valueStr, exists := c.Get(value)

	if !exists {
		return "", false
	}

	str, ok := valueStr.(string)
	if !ok {
		return "", false
	}

	return str, true

}

func GetUserIDFromContext(c *gin.Context) (string, error) {
	userID, exists := c.Get("user_id")
	if !exists {
		return "", errors.New("user ID not found in context")
	}

	userIDStr, ok := userID.(string)
	if !ok {
		return "", errors.New("user ID is not a string")
	}

	return userIDStr, nil
}

func ParseCustomTime(timeStr string, defaultTimezone *time.Location) (time.Time, error) {
	const layout = "2006-01-02 15:04:00"

	parsedTime, err := time.ParseInLocation(layout, timeStr, defaultTimezone)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
