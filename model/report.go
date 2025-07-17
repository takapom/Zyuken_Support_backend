package model

import "time"

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
