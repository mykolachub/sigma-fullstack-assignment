package middleware

import (
	"net/http"
	"sigma-test/config"
	"sigma-test/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func OnlyAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.Keys[config.PayloadUserRole]
		if role != config.AdminRole {
			message := util.MakeMessage(util.MessageError, config.ErrNoPermissions.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusForbidden, message)
			return
		}

		ctx.Next()
	}
}

func OnlyAdminOrOwner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payloadRole := ctx.Keys[config.PayloadUserRole]
		payloadId := ctx.Keys[config.PayloadUserId]

		isAdmin := payloadRole == config.AdminRole
		isOwner := payloadId == ctx.Query(config.QueryId)

		if !isOwner && !isAdmin {
			message := util.MakeMessage(util.MessageError, config.ErrNoPermissions.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusForbidden, message)
			return
		}

		ctx.Next()
	}
}

func Protect() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(config.AuthorizationHeader)
		if len(authHeader) == 0 {
			message := util.MakeMessage(util.MessageError, config.ErrNoAuthHeader.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, message)
			return
		}

		accessToken, err := util.ValidateBearerHeader(authHeader)
		if err != nil {
			message := util.MakeMessage(util.MessageError, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, message)
			return
		}

		token, err := util.ParseAndValidateJWTToken(accessToken)
		if err != nil {
			message := util.MakeMessage(util.MessageError, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, message)
			return
		}

		userId := token.Claims.(jwt.MapClaims)[config.JWTClaimsId]
		userRole := token.Claims.(jwt.MapClaims)[config.JWTClaimsRole]

		ctx.Set(config.PayloadUserRole, userRole)
		ctx.Set(config.PayloadUserId, userId)

		ctx.Next()
	}
}
