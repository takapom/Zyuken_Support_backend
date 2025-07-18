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

// ログインユーザー型定義
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// 新規登録のhandler処理
func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//上でreqに格納したものを取り出し、inputに入れている。サービス層に渡すための型変換
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

	//これはサービス層ではなく、クライアント層に返している
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

// ログイン処理のhandler処理
func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//サービス層向けに型を変更
	input := &service.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	}

	// サービス層を呼び出し、渡すのを試みる
	user, token, err := h.userService.Authenticate(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
		"token":   token,
	})
}

// RegisterRoutes ルートを登録する
func (h *UserHandler) RegisterRoutes(router *gin.Engine) {
	userRoutes := router.Group("/api/users")
	{
		userRoutes.POST("/register", h.Register)
		userRoutes.POST("/login", h.Login)
	}
}
