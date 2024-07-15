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

func (e *Usecase) Post() {
	e.sd.Post()
}
