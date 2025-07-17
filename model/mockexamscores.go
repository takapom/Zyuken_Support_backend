package model

import "time"

type MockExamScore struct {
	ID         string    `json:"id" db:"id"`
	MockExamID string    `json:"mock_exam_id" db:"mock_exam_id"`
	Subject    string    `json:"subject" db:"subject"` // 英語/数学/国語/理科/社会
	Score      float64   `json:"score" db:"score"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
