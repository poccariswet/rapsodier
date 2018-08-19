package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Task struct {
	Id        int       `gorm"id"`
	Text      string    `gorm:"text"`
	Done      bool      `gorm:"done"`
	Location  string    `gorm:"location"`
	CreatedAt time.Time `gorm:"created_at"`
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (t *Repository) SelectAll() ([]Task, error) {
	var tasks []Task
	if err := t.DB.Table(tablename).Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *Repository) Into(task Task) error {
	if err := t.DB.Table(tablename).Create(&task).Error; err != nil {
		return err
	}

	return nil
}

func (t *Repository) Update(task Task) error {
	if err := t.DB.Table(tablename).Where("id = ?", task.Id).Update("done", task.Done).Error; err != nil {
		return err
	}

	return nil
}
