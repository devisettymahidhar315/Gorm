package usecase

import (
	"app/adapaters"
	"app/core"
)

type Usecase struct {
	sd *adapaters.Database
}

func Init(s *adapaters.Database) *Usecase {
	return &Usecase{sd: s}
}

func (e *Usecase) Post(id int, name string) string {
	return e.sd.Post(id, name)
}

func (e *Usecase) Get_key(id int) (int, string) {
	return e.sd.Get_key(id)
}

func (e *Usecase) Get_Print() ([]core.User, string) {

	return e.sd.Get_Print()
}

func (e *Usecase) Del_key(id int) string {
	return e.sd.Del_key(id)
}

func (e *Usecase) Del_all() string {
	return e.sd.Del_all()
}
