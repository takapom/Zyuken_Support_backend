package model

import "time"

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
