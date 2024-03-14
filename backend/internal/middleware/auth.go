package middleware

import (
	"errors"
	"net/http"
	"sigma-test/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const authorizationHeader = "authorization"

const (
	payloadUserRole = "payload_user_role"
	payloadUserId   = "payload_user_id"
)

func OnlyAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.Keys[payloadUserRole]
		if role != "admin" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "you do not have permission to access this route"})
			return
		}

		ctx.Next()
	}
}

func OnlyAdminOrOwner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payloadRole := ctx.Keys[payloadUserRole]
		payloadId := ctx.Keys[payloadUserId]

		isAdmin := payloadRole == "admin"
		isOwner := payloadId == ctx.Query("id")

		if !isOwner && !isAdmin {
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

		accessToken, err := util.ValidateBearerHeader(authHeader)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		token, err := util.ParseAndValidateJWTToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		userId := token.Claims.(jwt.MapClaims)["id"]
		userRole := token.Claims.(jwt.MapClaims)["role"]

		ctx.Set(payloadUserRole, userRole)
		ctx.Set(payloadUserId, userId)

		ctx.Next()
	}
}
