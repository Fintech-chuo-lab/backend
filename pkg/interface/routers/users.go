package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/Fintech-chuo-lab/backend/pkg/models"
)

func ListUsers(c *gin.Context) {

	users := []models.User{
		{
			StudentID: "11c11111111",
			FirstName: "Ken",
			LastName: "Sue",
			Gender: 0,
			Age: 21,
			Grade: 4,
			Department: "商学",
		},
		{
			StudentID: "11c11111111",
			FirstName: "Ken",
			LastName: "Sue",
			Gender: 0,
			Age: 21,
			Grade: 4,
			Department: "商学",
		},
		{
			StudentID: "11c11111111",
			FirstName: "Ken",
			LastName: "Sue",
			Gender: 0,
			Age: 21,
			Grade: 4,
			Department: "商学",
		},
		{
			StudentID: "11c11111111",
			FirstName: "Ken",
			LastName: "Sue",
			Gender: 0,
			Age: 21,
			Grade: 4,
			Department: "商学",
		},
	}

	c.JSON(200, users)

}

func GetUser(c *gin.Context) {

	user := models.User{
		StudentID: "11c11111111",
		FirstName: "Ken",
		LastName: "Sue",
		Gender: 0,
		Age: 21,
		Grade: 4,
		Department: "商学",
	}

	c.JSON(200, user)

}
