package controller

import (
	"time"
	"users-rest/middleware"
	"users-rest/model"
	"users-rest/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

func (h *UserController) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if user.Role == "" {
		user.Role = "user"
	}

	createdUser, err := h.service.CreateUser(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	token, _ := h.GenerateToken(user.Name, user.Role)
	c.JSON(201, gin.H{"user": createdUser, "token": token})
}

func (h *UserController) GetUsers(c *gin.Context) {
	users, err := h.service.GetUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(200, users)
}

func (h *UserController) GenerateToken(username, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &model.Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(middleware.JwtKey)
}
