package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/takagiyuuki/zyuken-backend/middleware"
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

// 学校追加の型定義
type SchoolRequest struct {
	Name              string `json:"name" binding:"required"`
	Faculty           string `json:"faculty" binding:"required"`
	Level             string `json:"level" binding:"required,oneof=第一志望 併願 滑り止め"`
	ExamDate          string `json:"exam_date" binding:"required"` // フロントから文字列で受け取る
	Deviation         int    `json:"deviation" binding:"min=0,max=100"`
	PassRate          string `json:"pass_rate" binding:"omitempty,oneof=A B C D E"`
	ApplicationStatus string `json:"application_status" binding:"omitempty,oneof=出願予定 出願済 未出願"`
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

// ログイン処理のhandler処理
func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

// 新規学校の追加
func (h *UserHandler) NewSchool(c *gin.Context) {
	var req SchoolRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Contextからユーザー IDを取得（ミドルウェアで設定されたもの）
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	// ExamDateを文字列からtime.Timeに変換
	examDate, err := time.Parse("2006-01-02", req.ExamDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid exam date format"})
		return
	}

	//サービス層の型に変更
	input := &service.CreateSchoolInput{
		UserID:            userID.(string),
		Name:              req.Name,
		Faculty:           req.Faculty,
		Level:             req.Level,
		ExamDate:          examDate,
		Deviation:         req.Deviation,
		PassRate:          req.PassRate,
		ApplicationStatus: req.ApplicationStatus,
	}

	//サービス層に渡す
	school, err := h.userService.CreateSchool(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "School added successfully",
		"school":  school,
	})
}

// RegisterRoutes ルートを登録する
func (h *UserHandler) RegisterRoutes(router *gin.Engine) {
	userRoutes := router.Group("/api/users")
	{
		// 認証不要なエンドポイント
		userRoutes.POST("/register", h.Register)
		userRoutes.POST("/login", h.Login)

		// 認証が必要なエンドポイント
		auth := userRoutes.Group("/")
		auth.Use(middleware.AuthMiddleware())
		{
			auth.POST("/addschool", h.NewSchool)
		}
	}
}
