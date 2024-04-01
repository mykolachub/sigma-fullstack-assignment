package controller

import (
	"net/http"
	"sigma-test/config"
	"sigma-test/internal/middleware"
	"sigma-test/internal/request"
	"sigma-test/internal/util"

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

	r.GET("/api/user/me", middle.Protect(), handler.me)
	r.GET("/api/user", middle.Protect(), handler.getUserById)
	r.GET("/api/users", middle.Protect(), handler.getAllUsers)
	r.POST("/api/users", middle.Protect(), middle.OnlyAdmin(), handler.createUser)
	r.PATCH("/api/users", middle.Protect(), middle.OnlyAdminOrOwner(), handler.updateUser)
	r.DELETE("/api/users", middle.Protect(), middle.OnlyAdminOrOwner(), handler.deleteUser)
}

func (h UserHandler) me(c *gin.Context) {
	userId := c.Keys[config.PayloadUserId].(string)
	user, err := h.userSvc.GetUserById(userId)
	if err != nil {
		message := util.MakeMessage(util.MessageError, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusOK, util.MakeMessage(util.MessageSuccess, config.MsgEmpty, user))
}

func (h UserHandler) signup(c *gin.Context) {
	var body request.User
	if err := c.ShouldBindJSON(&body); err != nil {
		message := util.MakeMessage(util.MessageError, config.ErrFailedReadBody.Error(), nil)
		c.JSON(http.StatusBadRequest, message)
		return
	}

	user, err := h.userSvc.SignUp(body)
	if err != nil {
		message := util.MakeMessage(util.MessageError, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusOK, util.MakeMessage(util.MessageSuccess, config.MsgEmpty, user))
}

func (h UserHandler) login(c *gin.Context) {
	var body request.User
	if err := c.ShouldBindJSON(&body); err != nil {
		message := util.MakeMessage(util.MessageError, config.ErrFailedReadBody.Error(), nil)
		c.JSON(http.StatusBadRequest, message)
		return
	}

	token, err := h.userSvc.Login(body)
	if err != nil {
		message := util.MakeMessage(util.MessageError, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}

	c.JSON(http.StatusOK, util.MakeMessage(util.MessageSuccess, config.MsgEmpty, token))
}

func (h UserHandler) getAllUsers(c *gin.Context) {
	users, err := h.userSvc.GetAllUsers()
	if err != nil {
		message := util.MakeMessage(util.MessageError, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusOK, util.MakeMessage(util.MessageSuccess, config.MsgEmpty, users))
}

func (h UserHandler) getUserById(c *gin.Context) {
	id := c.Query(config.QueryId)
	if id == "" {
		message := util.MakeMessage(util.MessageError, config.ErrMissingIdPar.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}

	user, err := h.userSvc.GetUserById(id)
	if err != nil {
		message := util.MakeMessage(util.MessageError, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusOK, util.MakeMessage(util.MessageSuccess, config.MsgEmpty, user))
}

func (h UserHandler) createUser(c *gin.Context) {
	var body request.User
	if err := c.ShouldBindJSON(&body); err != nil {
		message := util.MakeMessage(util.MessageError, err.Error(), nil)
		c.JSON(http.StatusBadRequest, message)
		return
	}

	user, err := h.userSvc.CreateUser(body)
	if err != nil {
		message := util.MakeMessage(util.MessageError, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusCreated, util.MakeMessage(util.MessageSuccess, config.MsgUserCreated, user))
}

func (h UserHandler) updateUser(c *gin.Context) {
	id := c.Query(config.QueryId)
	if id == "" {
		message := util.MakeMessage(util.MessageError, config.ErrMissingIdPar.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	var body request.User

	if err := c.ShouldBindJSON(&body); err != nil {
		message := util.MakeMessage(util.MessageError, err.Error(), nil)
		c.JSON(http.StatusBadRequest, message)
		return
	}

	user, err := h.userSvc.UpdateUser(id, body)
	if err != nil {
		message := util.MakeMessage(util.MessageError, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusOK, util.MakeMessage(util.MessageSuccess, config.MsgUserUpdated, user))
}

func (h UserHandler) deleteUser(c *gin.Context) {
	id := c.Query(config.QueryId)
	if id == "" {
		message := util.MakeMessage(util.MessageError, config.ErrMissingIdPar.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}

	user, err := h.userSvc.DeleteUser(id)
	if err != nil {
		message := util.MakeMessage(util.MessageError, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusOK, util.MakeMessage(util.MessageSuccess, config.MsgUserDeleted, user))
}
