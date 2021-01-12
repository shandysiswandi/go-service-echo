package db

import (
	"go-rest-echo/db/mysql"
	"go-rest-echo/entity"
)

type db struct{}

// Interface is
type Interface interface {
	Start()
}

// NewDatabase is
func NewDatabase() Interface {
	return &db{}
}

func (db) Start() {
	err := mysql.NewDatabase().Initialize()
	if err != nil {
		panic(err)
	}
	mysql.GetDB().AutoMigrate(&entity.User{}, &entity.Task{})
}
