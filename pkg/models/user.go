package models

type User struct {
	StudentID string `json:"student_id"`
	NickName string `json:"nick_name"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Gender int `json:"gender"` // 0 is male, 1 is female
	Age int `json:"age"`
	Grade int `json:"grade"`
	Faculty string `json:"faculty"` // 学部
	Department string `json:"department"` //学科
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
