package model

import "time"

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
