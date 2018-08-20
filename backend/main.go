package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/tasks", ListTasks)
	router.POST("/tasks", NewTask)
	router.PUT("/tasks/:id", UpdateTask)
	router.Use(cors.Default())

	router.Run(":9999")
}
