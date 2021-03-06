package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Fintech-chuo-lab/backend/pkg/interface/routers"
    "net/http"
)

func init() {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	users := router.Group("/users")
	{
		users.GET("/", routers.ListUsers)
		users.GET("/:ID", routers.GetUser)
	}

	lessons := router.Group("/lessons")
	{
		lessons.GET("/", routers.ListLessons)
		lessons.GET("/:ID", routers.GetLesson)
	}

	http.Handle("/", router)
}
