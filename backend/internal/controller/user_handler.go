package controller

import (
	"net/http"
	"sigma-test/config"
	"sigma-test/internal/middleware"
	"sigma-test/internal/request"
	"sigma-test/internal/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandlerConfig struct {
	JwtSecret string
}

type UserHandler struct {
	userSvc UserService
	userCfg UserHandlerConfig
}

func InitUserHandler(r *gin.Engine, userSvc UserService, userCfg UserHandlerConfig) {
	handler := UserHandler{userSvc: userSvc, userCfg: userCfg}

	middle := middleware.InitMiddlewares(middleware.MiddlewareConfig{
		JwtSecret: handler.userCfg.JwtSecret,
	})

	r.POST("/api/user/signup", handler.signup)
	r.POST("/api/user/login", handler.login)

	users := r.Group("/api/users", middle.Protect())
	{
		users.GET("/me", handler.me)
		users.GET("/:user_id", handler.getUserById)
		users.GET("", handler.getAllUsers)
		users.POST("", middle.OnlyAdmin(), handler.createUser)
		users.PATCH("/:user_id", middle.OnlyAdminOrOwner(), handler.updateUser)
		users.DELETE("/:user_id", middle.OnlyAdminOrOwner(), handler.deleteUser)
	}
}

func (h UserHandler) me(c *gin.Context) {
	userId := c.Keys[config.PayloadUserId].(string)
	user, svcCode, err := h.userSvc.GetUserById(userId)
	if err != nil {
		message := util.NewErrResponse(svcCode.Message, svcCode.Code)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusOK, util.NewResponse(svcCode.Message, svcCode.Code).AddKey("user", user))
}

func (h UserHandler) signup(c *gin.Context) {
	var body request.User
	if err := c.ShouldBindJSON(&body); err != nil {
		svcCode := config.SvcFailedReadBody
		c.JSON(http.StatusBadRequest, util.NewErrResponse(svcCode.Message, svcCode.Code))
		return
	}

	user, svcCode, err := h.userSvc.SignUp(body)
	if err != nil {
		message := util.NewErrResponse(svcCode.Message, svcCode.Code)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusOK, util.NewResponse(svcCode.Message, svcCode.Code).AddKey("user", user))
}

func (h UserHandler) login(c *gin.Context) {
	var body request.User
	if err := c.ShouldBindJSON(&body); err != nil {
		svcCode := config.SvcFailedReadBody
		c.JSON(http.StatusBadRequest, util.NewErrResponse(svcCode.Message, svcCode.Code))
		return
	}

	token, svcCode, err := h.userSvc.Login(body)
	if err != nil {
		message := util.NewErrResponse(svcCode.Message, svcCode.Code)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}

	c.JSON(http.StatusOK, util.NewResponse(svcCode.Message, svcCode.Code).AddKey("token", token))
}

func (h UserHandler) getAllUsers(c *gin.Context) {
	search := c.Query(config.SearchParam)
	page, err := strconv.Atoi(c.Query(config.PageParam))
	if page < 0 || err != nil {
		page = 0
	}

	users, svcCode, err := h.userSvc.GetAllUsers(page, search)
	if err != nil {
		message := util.NewErrResponse(svcCode.Message, svcCode.Code)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusOK, util.NewResponse(svcCode.Message, svcCode.Code).AddKey("users", users))
}

func (h UserHandler) getUserById(c *gin.Context) {
	id := c.Param(config.UserId)
	if id == "" {
		svcCode := config.SvcMissingUserIdPar
		message := util.NewErrResponse(svcCode.Message, svcCode.Code)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}

	user, svcCode, err := h.userSvc.GetUserById(id)
	if err != nil {
		message := util.NewErrResponse(svcCode.Message, svcCode.Code)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusOK, util.NewResponse(svcCode.Message, svcCode.Code).AddKey("user", user))
}

func (h UserHandler) createUser(c *gin.Context) {
	var body request.User
	if err := c.ShouldBindJSON(&body); err != nil {
		svcCode := config.SvcFailedReadBody
		c.JSON(http.StatusBadRequest, util.NewErrResponse(svcCode.Message, svcCode.Code))
		return
	}

	user, svcCode, err := h.userSvc.CreateUser(body)
	if err != nil {
		message := util.NewErrResponse(svcCode.Message, svcCode.Code)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusCreated, util.NewResponse(svcCode.Message, svcCode.Code).AddKey("user", user))
}

func (h UserHandler) updateUser(c *gin.Context) {
	id := c.Param(config.UserId)
	if id == "" {
		svcCode := config.SvcMissingUserIdPar
		message := util.NewErrResponse(svcCode.Message, svcCode.Code)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	var body request.User

	if err := c.ShouldBindJSON(&body); err != nil {
		svcCode := config.SvcFailedReadBody
		c.JSON(http.StatusBadRequest, util.NewErrResponse(svcCode.Message, svcCode.Code))
		return
	}

	user, svcCode, err := h.userSvc.UpdateUser(id, body)
	if err != nil {
		message := util.NewErrResponse(svcCode.Message, svcCode.Code)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusOK, util.NewResponse(svcCode.Message, svcCode.Code).AddKey("user", user))
}

func (h UserHandler) deleteUser(c *gin.Context) {
	id := c.Param(config.UserId)
	if id == "" {
		svcCode := config.SvcMissingUserIdPar
		message := util.NewErrResponse(svcCode.Message, svcCode.Code)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}

	user, svcCode, err := h.userSvc.DeleteUser(id)
	if err != nil {
		message := util.NewErrResponse(svcCode.Message, svcCode.Code)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusOK, util.NewResponse(svcCode.Message, svcCode.Code).AddKey("user", user))
}
