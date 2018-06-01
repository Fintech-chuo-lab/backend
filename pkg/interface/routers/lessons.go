package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/Fintech-chuo-lab/backend/pkg/models"
)

func ListLessons(c *gin.Context) {

	lessons := []models.Lesson{
		{
			Name: "日本経済論",
			Teacher: "田中？",
			Credit: 4,
			Classroom: "8305",
			Department: "商学",
			CanTakeLessonGrade: []int{3,4},
		},
		{
			Name: "日本経済論",
			Teacher: "田中？",
			Credit: 4,
			Classroom: "8305",
			Department: "商学",
			CanTakeLessonGrade: []int{3,4},
		},
		{
			Name: "日本経済論",
			Teacher: "田中？",
			Credit: 4,
			Classroom: "8305",
			Department: "商学",
			CanTakeLessonGrade: []int{3,4},
		},
		{
			Name: "日本経済論",
			Teacher: "田中？",
			Credit: 4,
			Classroom: "8305",
			Department: "商学",
			CanTakeLessonGrade: []int{3,4},
		},
	}

	c.JSON(200, lessons)

}

func GetLesson(c *gin.Context) {

	lesson := models.Lesson{
		Name: "日本経済論",
		Teacher: "田中？",
		Credit: 4,
		Classroom: "8305",
		Department: "商学",
		CanTakeLessonGrade: []int{3,4},
	}

	c.JSON(200, lesson)

}
