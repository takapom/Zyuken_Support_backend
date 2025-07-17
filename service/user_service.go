// ユーザーに関する処理
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

type CreateUserInput struct {
	Email          string
	Password       string
	Name           string
	Department     string
	GraduationYear int
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
