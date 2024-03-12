package controller

import (
	"net/http"
	"sigma-test/internal/middleware"
	"sigma-test/internal/request"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userSvc UserService
}

func InitUserHandler(r *gin.Engine, userSvc UserService) {
	handler := UserHandler{userSvc: userSvc}

	r.POST("/api/signup", handler.signup)
	r.POST("/api/login", handler.login)

	r.Use(middleware.Protect())

	r.GET("/api/me", handler.me)

	r.GET("/api/user", handler.getUserById)
	r.GET("/api/users", handler.getAllUsers)
	r.POST("/api/users", middleware.OnlyAdmin(), handler.createUser)
	r.PATCH("/api/users", middleware.OnlyAdminOrOwner(), handler.updateUser)
	r.DELETE("/api/users", middleware.OnlyAdminOrOwner(), handler.deleteUser)

}

func (h UserHandler) me(c *gin.Context) {
	userId := c.Keys["payload_user_id"].(string)
	user, err := h.userSvc.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h UserHandler) signup(c *gin.Context) {
	var body request.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	user, err := h.userSvc.SignUp(body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h UserHandler) login(c *gin.Context) {
	var body request.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	token, err := h.userSvc.Login(body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h UserHandler) getAllUsers(c *gin.Context) {
	users, err := h.userSvc.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (h UserHandler) getUserById(c *gin.Context) {
	id := c.Query("id")
	user, err := h.userSvc.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h UserHandler) createUser(c *gin.Context) {
	var body request.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userSvc.CreateUser(body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h UserHandler) updateUser(c *gin.Context) {
	id := c.Query("id")
	var body request.User

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userSvc.UpdateUser(id, body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h UserHandler) deleteUser(c *gin.Context) {
	id := c.Query("id")
	err := h.userSvc.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User successfully deleted"})
}
