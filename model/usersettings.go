package model

import "time"

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
