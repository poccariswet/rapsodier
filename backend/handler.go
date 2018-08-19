package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTasks(c *gin.Context) {
	repo, err := NewConfig()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"config error": err})
	}
	defer repo.DB.Close()

	tasks, err := repo.SelectAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"select error": err})
	}

	c.JSON(http.StatusOK, tasks)
}

func NewTask(c *gin.Context) {
	repo, err := NewConfig()
	if err != nil {
		log.Println("config err: ", err)
	}
	defer repo.DB.Close()

	var task Task
	task.Location = "todo" // default location value

	if err := c.Bind(&task); err != nil {
		log.Println("bind err: ", err)
	}

	if err := repo.Into(task); err != nil {
		log.Println("into err: ", err)
	}
}

func UpdateTask(c *gin.Context) {
	repo, err := NewConfig()
	if err != nil {
		log.Println("config err: ", err)
	}
	defer repo.DB.Close()

	var task Task
	if err := c.Bind(&task); err != nil {
		log.Println("bind err: ", err)
	}

	if err := repo.Update(task); err != nil {
		log.Println("update err: ", err)
	}
}
