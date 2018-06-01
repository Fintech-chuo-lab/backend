package models

type Lesson struct {
	Name string
	Teacher string
	Credit int
	Classroom string
	Department string
	CanTakeLessonGrade []int
}
