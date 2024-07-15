package api

import "app/usecase"

type Api struct {
	as *usecase.Usecase
}

func Init(a *usecase.Usecase) *Api {
	return &Api{as: a}
}

func (a *Api) Post() {
	a.as.Post()
}
