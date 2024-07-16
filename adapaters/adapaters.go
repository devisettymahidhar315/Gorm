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

func (e *Database) Post(id int, name string) string {

	e1 := core.User{Id: id, Name: name}

	result := e.df.Create(&e1)

	if result.RowsAffected == 1 {
		return "data is added succesfully"
	} else {

		return result.Error.Error()
	}

}

func (e *Database) Get_key(id int) (int, string) {

	var e1 core.User

	result := e.df.First(&e1, id)
	if result.Error != nil {
		return 0, result.Error.Error()
	} else {
		return e1.Id, e1.Name
	}

}

func (e *Database) Get_Print() ([]core.User, string) {
	var e1 []core.User
	result := e.df.Find(&e1)
	if result.Error != nil {
		return nil, result.Error.Error()
	} else {
		return e1, "nil"
	}

}

func (e *Database) Del_key(id int) string {
	var e1 *core.User
	result := e.df.Delete(&e1, "Id = ?", id)
	if result.Error != nil {
		return result.Error.Error()
	} else if result.RowsAffected == 0 {
		return "key not found"
	} else {
		return "key is deleted"
	}
}

func (e *Database) Del_all() string {

	e.df.Unscoped().Where("1 = 1").Delete(&core.User{})
	e.df.Exec("ALTER TABLE users AUTO_INCREMENT = 1")
	return "deleted all rows in the table"
}
