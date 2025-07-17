package model

import "time"

type MockExam struct {
	ID           string    `json:"id" db:"id"`
	UserID       string    `json:"user_id" db:"user_id"`
	Name         string    `json:"name" db:"name"`
	Date         time.Time `json:"date" db:"date"`
	Status       string    `json:"status" db:"status"` // 受験予定/結果返却
	OverallScore float64   `json:"overall_score" db:"overall_score"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
