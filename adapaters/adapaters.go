package adapaters

import (
	"app/core"
	"fmt"

	"gorm.io/gorm"
)

type Database struct {
	df *gorm.DB
}

func Init(db *gorm.DB) *Database {
	return &Database{df: db}

}

func (e *Database) Post(id int, name string) {

	e1 := core.User{Id: id, Name: name}

	e.df.Create(&e1)
}

func (e *Database) Get_key(id int) {
	var e1 core.User

	e.df.Find(&e1, "Id = ?", id)
	fmt.Println(e1)
}

func (e *Database) Get_Print() {
	var e1 []core.User
	e.df.Find(&e1)
	fmt.Println(e1)
}

func (e *Database) Del_key(id int) {
	var e1 *core.User
	e.df.Delete(&e1, "Id = ?", id)
}

func (e *Database) Del_all() {

	e.df.Unscoped().Where("1 = 1").Delete(&core.User{})
	e.df.Exec("ALTER TABLE users AUTO_INCREMENT = 1")

}
