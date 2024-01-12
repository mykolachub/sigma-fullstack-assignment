package controller

import (
	"net/http"
	"sigma-test/internal/request"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userSvc UserService
}

func InitUserHandler(r *gin.Engine, userSvc UserService) {
	handler := UserHandler{userSvc: userSvc}

	// Basic CRUD endpoint
	r.GET("/api/user", handler.getUserById)
	r.GET("/api/users", handler.getAllUsers)
	r.POST("/api/users", handler.createUser)
	r.PATCH("/api/users", handler.updateUser)
	r.DELETE("/api/users", handler.deleteUser)
}

func (h UserHandler) getAllUsers(c *gin.Context) {
	users, err := h.userSvc.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (h UserHandler) getUserById(c *gin.Context) {
	id := c.Query("id")
	user, err := h.userSvc.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h UserHandler) createUser(c *gin.Context) {
	var body request.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := h.userSvc.CreateUser(body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h UserHandler) updateUser(c *gin.Context) {
	id := c.Query("id")
	var body request.User

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := h.userSvc.UpdateUser(id, body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h UserHandler) deleteUser(c *gin.Context) {
	id := c.Query("id")
	err := h.userSvc.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User successfully deleted"})
}
