package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTasks(c *gin.Context) {
	repo, err := NewConfig()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	defer repo.DB.Close()

	tasks, err := repo.SelectAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, tasks)
}

func NewTask(c *gin.Context) {
	repo, err := NewConfig()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	defer repo.DB.Close()

	var task Task
	if err := c.Bind(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	if err := repo.Into(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func UpdateTask(c *gin.Context) {
	repo, err := NewConfig()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	defer repo.DB.Close()

	var task Task
	if err := c.Bind(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	if err := repo.Update(task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
