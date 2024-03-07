package controller

import (
	"fmt"
	"net/http"
	"sigma-test/internal/request"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

	r.POST("/api/signup", handler.signup)
	r.POST("/api/login", handler.login)

}

func (h UserHandler) signup(c *gin.Context) {
	// get email/password/role from the body
	var body request.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	// chech if the user already exists
	u, err := h.userSvc.GetUserByEmail(body.Email)
	fmt.Println(u, err)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash body"})
		return
	}

	// create user
	user, err := h.userSvc.CreateUser(request.User{Email: body.Email, Password: string(hash), Role: body.Role})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Failed to create user"})
		return
	}

	// respond
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h UserHandler) login(c *gin.Context) {
	// get the email, password off req body
	var body request.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	// look up requested user
	user, err := h.userSvc.GetUserByEmail(body.Email)
	fmt.Println(user, err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "1. Invalid email or password"})
		return
	}

	// compare sent in pass with saved user pass hash
	fmt.Println(user.Password, body.Password, user.Password == body.Password)
	hash_err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if hash_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": hash_err.Error()})
		return
	}

	// generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create token"})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// send it back to client
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
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
