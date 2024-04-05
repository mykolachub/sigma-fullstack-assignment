package middleware

import (
	"net/http"
	"sigma-test/config"
	"sigma-test/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type MiddlewareConfig struct {
	JwtSecret string
}

type Middleware struct {
	cfg MiddlewareConfig
}

func InitMiddlewares(cfg MiddlewareConfig) Middleware {
	return Middleware{cfg: cfg}
}

func (m Middleware) OnlyAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.Keys[config.PayloadUserRole]
		if role != config.AdminRole {
			svcCode := config.SvcNoPermissions
			message := util.NewErrResponse(svcCode.Message, svcCode.Code)
			ctx.AbortWithStatusJSON(http.StatusForbidden, message)
			return
		}

		ctx.Next()
	}
}

func (m Middleware) OnlyAdminOrOwner() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payloadRole := ctx.Keys[config.PayloadUserRole]
		payloadId := ctx.Keys[config.PayloadUserId]

		isAdmin := payloadRole == config.AdminRole
		isOwner := payloadId == ctx.Param(config.UserId)

		if !isOwner && !isAdmin {
			svcCode := config.SvcNoPermissions
			message := util.NewErrResponse(svcCode.Message, svcCode.Code)
			ctx.AbortWithStatusJSON(http.StatusForbidden, message)
			ctx.AbortWithStatusJSON(http.StatusForbidden, message)
			return
		}

		ctx.Next()
	}
}

func (m Middleware) Protect() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(config.AuthorizationHeader)
		if len(authHeader) == 0 {
			svcCode := config.SvcNoAuthHeader
			message := util.NewErrResponse(svcCode.Message, svcCode.Code)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, message)
			return
		}

		accessToken, svcCode, err := util.ValidateBearerHeader(authHeader)
		if err != nil {
			message := util.NewErrResponse(svcCode.Message, svcCode.Code)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, message)
			return
		}

		token, svcCode, err := util.ParseAndValidateJWTToken(accessToken, m.cfg.JwtSecret)
		if err != nil {
			message := util.NewErrResponse(svcCode.Message, svcCode.Code)
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
