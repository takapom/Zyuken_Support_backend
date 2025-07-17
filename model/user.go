package model

import "time"

// 1. Users（ユーザー）
type User struct {
	ID             string    `json:"id" db:"id"`
	Email          string    `json:"email" db:"email"`
	Password       string    `json:"password" db:"password"`
	Name           string    `json:"name" db:"name"`
	Department     string    `json:"department" db:"department"` // 理系/文系/医系
	GraduationYear int       `json:"graduation_year" db:"graduation_year"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

//   5. Costs（費用）

type Cost struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	Category  string    `json:"category" db:"category"` // 受験料/交通費/宿泊費/教材費
	Name      string    `json:"name" db:"name"`
	Amount    int       `json:"amount" db:"amount"`
	Date      time.Time `json:"date" db:"date"`
	Status    string    `json:"status" db:"status"` // 支払い済/未払い/予定/予約済
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

//   6. Tasks（タスク）

type Task struct {
	ID        string     `json:"id" db:"id"`
	UserID    string     `json:"user_id" db:"user_id"`
	Title     string     `json:"title" db:"title"`
	Completed bool       `json:"completed" db:"completed"`
	DueDate   *time.Time `json:"due_date" db:"due_date"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
}

//   7. Schedules（スケジュール）

type Schedule struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	Title     string    `json:"title" db:"title"`
	Date      time.Time `json:"date" db:"date"`
	Type      string    `json:"type" db:"type"` // 模試/入試/その他
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

//   8. Reports（レポート）

type Report struct {
	ID          string    `json:"id" db:"id"`
	UserID      string    `json:"user_id" db:"user_id"`
	Type        string    `json:"type" db:"type"` // 成績推移/受験校分析/費用サマリー/学習計画
	FileName    string    `json:"file_name" db:"file_name"`
	FileSize    string    `json:"file_size" db:"file_size"`
	FilePath    string    `json:"file_path" db:"file_path"`
	GeneratedAt time.Time `json:"generated_at" db:"generated_at"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

//   9. UserSettings（ユーザー設定）

type UserSettings struct {
	ID                      string `json:"id" db:"id"`
	UserID                  string `json:"user_id" db:"user_id"`
	ExamReminder            bool   `json:"exam_reminder" db:"exam_reminder"`
	ApplicationDeadline     bool   `json:"application_deadline" db:"application_deadline"`
	ScoreUpdateNotification bool   `json:"score_update_notification" 
  db:"score_update_notification"`
	DataPrivacy string `json:"data_privacy" db:"data_privacy"` //
	//   非公開/保護者と共有/先生と共有
	TwoFactorAuth bool      `json:"two_factor_auth" db:"two_factor_auth"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

//   10. Budget（予算）

type Budget struct {
	ID          string    `json:"id" db:"id"`
	UserID      string    `json:"user_id" db:"user_id"`
	TotalAmount int       `json:"total_amount" db:"total_amount"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
