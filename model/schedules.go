package model

import "time"

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
