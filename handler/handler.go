package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takagiyuuki/zyuken-backend/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// 新規登録ユーザー型定義
type RegisterRequest struct {
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required,min=6"`
	Name           string `json:"name" binding:"required"`
	Department     string `json:"department" binding:"required"`
	GraduationYear int    `json:"graduation_year" binding:"required"`
}

func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	input := &service.CreateUserInput{
		Email:          req.Email,
		Password:       req.Password,
		Name:           req.Name,
		Department:     req.Department,
		GraduationYear: req.GraduationYear,
	}

	user, err := h.userService.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

func (h *UserHandler) RegisterRoutes(router *gin.Engine) {
	userRoutes := router.Group("/api/users")
	{
		userRoutes.POST("/register", h.Register)
	}
}
