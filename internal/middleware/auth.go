package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lifeneedzoya/habit-tracker/internal/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string

		token, err := c.Cookie("Authorization")

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			c.Abort()
			return
		}

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			c.Abort()
			return
		}

		claims, err := utils.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		if userID, ok := (*claims)["user_id"].(string); ok {
			c.Set("user_id", userID)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Next()
	}
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

func AuthMiddlewareWithConfig(config AuthConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string

	tokenLoop:
		for _, source := range config.TokenSources {
			switch source {
			case "header":
				if bearerToken := c.GetHeader("Authorization"); bearerToken != "" {
					if tokenFromHeader, hasPrefix := strings.CutPrefix(bearerToken, "Bearer "); hasPrefix {
						token = tokenFromHeader
						break tokenLoop
					}
				}
			case "cookie":
				if cookieToken, err := c.Cookie(config.CookieName); err == nil {
					token = cookieToken
					break tokenLoop
				}
			case "query":
				if queryToken := c.Query("token"); queryToken != "" {
					token = queryToken
					break tokenLoop
				}
			}
		}

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": config.ErrorMessage})
			c.Abort()
			return
		}

		claims, err := utils.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		if userID, ok := (*claims)["user_id"].(string); ok {
			c.Set("user_id", userID)
		}

		c.Set("token_claims", claims)

		c.Next()
	}
}

type AuthConfig struct {
	TokenSources []string // ["header", "cookie", "query"]
	CookieName   string
	ErrorMessage string
}

func DefaultAuthConfig() AuthConfig {
	return AuthConfig{
		TokenSources: []string{"header", "cookie"},
		CookieName:   "Authorization",
		ErrorMessage: "Authentication required",
	}
}
