package usecase

import (
	"app/adapaters"
)

type Usecase struct {
	sd *adapaters.Database
}

func Init(s *adapaters.Database) *Usecase {
	return &Usecase{sd: s}
}

func (e *Usecase) Post(id int, name string) {
	e.sd.Post(id, name)
}

func (e *Usecase) Get_key(id int) {
	e.sd.Get_key(id)
}

func (e *Usecase) Get_Print() {
	e.sd.Get_Print()
}

func (e *Usecase) Del_key(id int) {
	e.sd.Del_key(id)
}

func (e *Usecase) Del_all() {
	e.sd.Del_all()
}
