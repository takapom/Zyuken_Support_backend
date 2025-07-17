package model

import "time"

type School struct {
	ID                string    `json:"id" db:"id"`
	UserID            string    `json:"user_id" db:"user_id"`
	Name              string    `json:"name" db:"name"`
	Faculty           string    `json:"faculty" db:"faculty"`
	Level             string    `json:"level" db:"level"` // 第一志望/併願/滑り止め
	ExamDate          time.Time `json:"exam_date" db:"exam_date"`
	Deviation         int       `json:"deviation" db:"deviation"`
	PassRate          string    `json:"pass_rate" db:"pass_rate"`                   // A/B/C/D/E判定
	ApplicationStatus string    `json:"application_status" db:"application_status"` //
	//   出願予定/出願済/未出願
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
