package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

type CreateSchoolInput struct {
	UserID            string
	Name              string
	Faculty           string
	Level             string
	ExamDate          time.Time
	Deviation         int
	PassRate          string
	ApplicationStatus string
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

// サービス層のログイン認証
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

	//comparehashandpasswordはdbと引数のパスワードを比較して検証
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	// JWTトークンの生成
	token, err := generateJWT(user.ID)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	// パスワードを削除してから返す
	user.Password = ""
	return user, token, nil
}

// 新規学校の追加処理
func (s *UserService) CreateSchool(newschool *CreateSchoolInput) (*model.School, error) {
	if newschool.UserID == "" {
		return nil, errors.New("user_id is required")
	}
	if newschool.Name == "" {
		return nil, errors.New("school name is required")
	}
	if newschool.Faculty == "" {
		return nil, errors.New("faculty is required")
	}
	if newschool.Level == "" {
		return nil, errors.New("level is required")
	}
	if newschool.ExamDate.IsZero() {
		return nil, errors.New("exam date is required")
	}
	if newschool.Deviation < 0 || newschool.Deviation > 100 {
		return nil, errors.New("deviation must be between 0 and 100")
	}
	if newschool.ApplicationStatus != "" {
		validStatuses := map[string]bool{
			"出願予定": true,
			"出願済":  true,
			"未出願":  true,
		}
		if !validStatuses[newschool.ApplicationStatus] {
			return nil, errors.New("invalid application status")
		}
	}

	school := &model.School{
		ID:                uuid.New().String(),
		UserID:            newschool.UserID,
		Name:              newschool.Name,
		Faculty:           newschool.Faculty,
		Level:             newschool.Level,
		ExamDate:          newschool.ExamDate,
		Deviation:         newschool.Deviation,
		PassRate:          newschool.PassRate,
		ApplicationStatus: newschool.ApplicationStatus,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	if err := s.userRepo.AddSchool(school); err != nil {
		return nil, err
	}

	return school, nil
}

// JWT生成関数
func generateJWT(userID string) (string, error) {
	// 秘密鍵（本番環境では環境変数から取得すべき）
	secretKey := []byte("your-secret-key-here")

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
		"iat":     time.Now().Unix(),
	}

	// トークンを作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 秘密鍵で署名してトークン文字列を生成
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
