package controller

import (
	"fmt"
	"net/http"
	"os"
	"sigma-test/internal/middleware"
	"sigma-test/internal/request"
	"sigma-test/internal/response"
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

	r.POST("/api/signup", handler.signup)
	r.POST("/api/login", handler.login)

	r.Use(middleware.Protect(userSvc))

	r.GET("/api/me", handler.me)

	r.GET("/api/user", handler.getUserById)
	r.GET("/api/users", handler.getAllUsers)
	r.POST("/api/users", middleware.OnlyAdmin(), handler.createUser)
	r.PATCH("/api/users", middleware.OnlyAdminOrOwner(), handler.updateUser)
	r.DELETE("/api/users", middleware.OnlyAdminOrOwner(), handler.deleteUser)
}

func (h UserHandler) me(c *gin.Context) {
	payloadId := c.Keys["payload_user_id"] // create getters in auth.go and user them instead of direct access
	payloadEmail := c.Keys["payload_user_email"]
	payloadRole := c.Keys["payload_user_role"]
	payloadPassword := c.Keys["payload_user_password"]

	c.JSON(http.StatusOK, gin.H{"data": response.User{
		ID:       payloadId.(string),
		Email:    payloadEmail.(string),
		Password: payloadPassword.(string),
		Role:     payloadRole.(string),
	}})
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
	user, err := h.userSvc.CreateUser(
		request.User{Email: body.Email, Password: string(hash), Role: body.Role},
	)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	// compare sent in pass with saved user pass hash
	fmt.Println(user.Password, body.Password, user.Password == body.Password)
	// best to move this to a service and use strategy pattern for different hashing algorithms (bcrypt, scrypt, argon2) in case of future changes
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

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
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
