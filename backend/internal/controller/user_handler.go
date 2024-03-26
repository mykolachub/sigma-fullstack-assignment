package controller

import (
	"net/http"
	"sigma-test/internal/middleware"
	"sigma-test/internal/request"
	"sigma-test/internal/util"

	"github.com/gin-gonic/gin"
)

const playloadUserID = "payload_user_id"

type UserHandler struct {
	userSvc UserService
}

func InitUserHandler(r *gin.Engine, userSvc UserService) {
	handler := UserHandler{userSvc: userSvc}

	r.POST("/api/user/signup", handler.signup)
	r.POST("/api/user/login", handler.login)

	r.Use(middleware.Protect())

	r.GET("/api/user/me", handler.me)

	r.GET("/api/user", handler.getUserById)
	r.GET("/api/users", handler.getAllUsers)
	r.POST("/api/users", middleware.OnlyAdmin(), handler.createUser)
	r.PATCH("/api/users", middleware.OnlyAdminOrOwner(), handler.updateUser)
	r.DELETE("/api/users", middleware.OnlyAdminOrOwner(), handler.deleteUser)

}

func (h UserHandler) me(c *gin.Context) {
	userId := c.Keys[playloadUserID].(string)
	user, err := h.userSvc.GetUserById(userId)
	if err != nil {
		message := util.MakeMessage(util.MessageError, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusOK, util.MakeMessage(util.MessageSuccess, "", user))
}

func (h UserHandler) signup(c *gin.Context) {
	var body request.User
	if err := c.ShouldBindJSON(&body); err != nil {
		message := util.MakeMessage(util.MessageError, "failed to read body", nil)
		c.JSON(http.StatusBadRequest, message)
		return
	}

	user, err := h.userSvc.SignUp(body)
	if err != nil {
		message := util.MakeMessage(util.MessageError, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusOK, util.MakeMessage(util.MessageSuccess, "", user))
}

func (h UserHandler) login(c *gin.Context) {
	var body request.User
	if err := c.ShouldBindJSON(&body); err != nil {
		message := util.MakeMessage(util.MessageError, "failed to read body", nil)
		c.JSON(http.StatusBadRequest, message)
		return
	}

	token, err := h.userSvc.Login(body)
	if err != nil {
		message := util.MakeMessage(util.MessageError, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}

	c.JSON(http.StatusOK, util.MakeMessage(util.MessageSuccess, "", token))
}

func (h UserHandler) getAllUsers(c *gin.Context) {
	users, err := h.userSvc.GetAllUsers()
	if err != nil {
		message := util.MakeMessage(util.MessageError, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusOK, util.MakeMessage(util.MessageSuccess, "", users))
}

func (h UserHandler) getUserById(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		message := util.MakeMessage(util.MessageError, "missing id parameter", nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}

	user, err := h.userSvc.GetUserById(id)
	if err != nil {
		message := util.MakeMessage(util.MessageError, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusOK, util.MakeMessage(util.MessageSuccess, "", user))
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
	c.JSON(http.StatusCreated, util.MakeMessage(util.MessageSuccess, "user created", user))
}

func (h UserHandler) updateUser(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		message := util.MakeMessage(util.MessageError, "missing id parameter", nil)
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
	c.JSON(http.StatusOK, util.MakeMessage(util.MessageSuccess, "user updated", user))
}

func (h UserHandler) deleteUser(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		message := util.MakeMessage(util.MessageError, "missing id parameter", nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}

	user, err := h.userSvc.DeleteUser(id)
	if err != nil {
		message := util.MakeMessage(util.MessageError, err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, message)
		return
	}
	c.JSON(http.StatusOK, util.MakeMessage(util.MessageSuccess, "user deleted", user))
}
