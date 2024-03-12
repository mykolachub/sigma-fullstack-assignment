package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"sigma-test/internal/response"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	authorizationHeader     = "authorization"
	authorizationTypeBearer = "bearer"
)

type ContextKey string // separate type for context keys to avoid clashin

const playoadUserRole ContextKey = "payload_user_role"

// depdend only on what it needs
type UserService interface {
	GetUserById(id string) (response.User, error)
}

func OnlyAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.Keys["payload_user_role"]
		if role != "admin" {
			ctx.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{"error": "you do not have permission to access this route"},
			)
			return
		}

		ctx.Next()
	}
}

func OnlyAdminOrOwner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payloadRole := ctx.Keys["payload_user_role"] // const
		payloadId := ctx.Keys["payload_user_id"]     // const

		isAdmin := payloadRole == "admin"       // const
		isOwner := payloadId == ctx.Query("id") // const

		if !isOwner && !isAdmin { // [NIT] De Morgan's law
			ctx.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{"error": "you do not have permission to access this route"},
			)
			return
		}

		ctx.Next()
	}
}

func Protect(userService UserService) gin.HandlerFunc {
	// Split into multile functions: validate, parse token
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(authorizationHeader)
		if len(authHeader) == 0 {
			err := errors.New("authorization header not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		authType := strings.ToLower(fields[0])
		if authType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		accessToken := fields[1]
		token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if !token.Valid {
			err := fmt.Errorf("invalid token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		userId, err := token.Claims.GetSubject()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		payload, err := userService.GetUserById(userId)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("payload_user_role", payload.Role) // package level consts e.g. playoad_user_role
		ctx.Set("payload_user_id", payload.ID)
		ctx.Set("payload_user_email", payload.Email)
		ctx.Set("payload_user_password", payload.Password)

		ctx.Next()
	}
}