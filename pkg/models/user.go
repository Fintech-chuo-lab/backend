package models

type User struct {
	StudentID string `json:"student_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Gender int `json:"gender"` // 0 is male, 1 is female
	Age int `json:"age"`
	Grade int `json:"grade"`
	Department string `json:"department"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
