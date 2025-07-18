package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/takagiyuuki/zyuken-backend/model"
	"github.com/takagiyuuki/zyuken-backend/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// 新規登録用の入力型
type CreateUserInput struct {
	Email          string
	Password       string
	Name           string
	Department     string
	GraduationYear int
}

// ログイン用の入力型
type LoginInput struct {
	Email    string
	Password string
}

func (s *UserService) CreateUser(input *CreateUserInput) (*model.User, error) {
	if input.Email == "" {
		return nil, errors.New("email is required")
	}
	if input.Password == "" {
		return nil, errors.New("password is required")
	}
	if input.Name == "" {
		return nil, errors.New("name is required")
	}
	if input.Department == "" {
		return nil, errors.New("department is required")
	}
	if input.GraduationYear == 0 {
		return nil, errors.New("graduation year is required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:             uuid.New().String(),
		Email:          input.Email,
		Password:       string(hashedPassword),
		Name:           input.Name,
		Department:     input.Department,
		GraduationYear: input.GraduationYear,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}

func (s *UserService) Authenticate(input *LoginInput) (*model.User, string, error) {
	if input.Email == "" {
		return nil, "", errors.New("email is required")
	}
	if input.Password == "" {
		return nil, "", errors.New("password is required")
	}

	// メールアドレスでユーザーを検索
	user, err := s.userRepo.FindByEmail(input.Email)
	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	// パスワードの検証
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	// TODO: JWTトークンの生成（現在は仮のトークン）
	token := "dummy-token-" + user.ID

	// パスワードを削除してから返す
	user.Password = ""
	return user, token, nil
}
