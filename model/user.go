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
