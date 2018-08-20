package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	DB *gorm.DB
}

const tablename = "todos"

var (
	user   = os.Getenv("MYSQL_USER")
	pass   = os.Getenv("MYSQL_PASS")
	dbname = os.Getenv("MYSQL_DATABASE")
)

func DBconnect() (*gorm.DB, error) {
	connection := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, dbname)
	db, err := gorm.Open("mysql", connection)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewConfig() (*Repository, error) {
	db, err := DBconnect()
	if err != nil {
		return nil, err
	}
	repo := NewRepository(db)

	return repo, nil
}
