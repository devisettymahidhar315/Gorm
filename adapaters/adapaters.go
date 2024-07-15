package adapaters

import (
	"app/core"

	"gorm.io/gorm"
)

type Database struct {
	df *gorm.DB
}

func Init(db *gorm.DB) *Database {
	return &Database{df: db}

}

func (e *Database) Post() {

	e1 := core.User{Id: 1, Name: "kl"}

	e.df.Create(&e1)
}
