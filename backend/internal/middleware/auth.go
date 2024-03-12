package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	authorizationHeader     = "authorization"
	authorizationTypeBearer = "bearer"
)

func OnlyAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.Keys["payload_user_role"]
		if role != "admin" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "you do not have permission to access this route"})
			return
		}

		ctx.Next()
	}
}

func OnlyAdminOrOwner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payloadRole := ctx.Keys["payload_user_role"]
		payloadId := ctx.Keys["payload_user_id"]

		isAdmin := payloadRole == "admin"
		isOwner := payloadId == ctx.Query("id")

		if !(isOwner || isAdmin) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "you do not have permission to access this route"})
			return
		}

		ctx.Next()
	}
}

func Protect() gin.HandlerFunc {
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

		userId := token.Claims.(jwt.MapClaims)["id"]
		userRole := token.Claims.(jwt.MapClaims)["role"]

		ctx.Set("payload_user_role", userRole)
		ctx.Set("payload_user_id", userId)

		ctx.Next()
	}
}
