package api

import (
	"app/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Api struct {
	as *usecase.Usecase
}

func Init(a *usecase.Usecase) *Api {
	return &Api{as: a}
}

func (a *Api) Post(ctx *gin.Context) {
	i := ctx.Param("id")
	name := ctx.Param("name")
	id, _ := strconv.Atoi(i)
	a.as.Post(id, name)
}

func (a *Api) Get_key(ctx *gin.Context) {
	i := ctx.Param("id")
	id, _ := strconv.Atoi(i)
	a.as.Get_key(id)
}

func (a *Api) Get_Print(ctx *gin.Context) {
	a.as.Get_Print()
}

func (a *Api) Del_all(ctx *gin.Context) {
	a.as.Del_all()
}

func (a *Api) Del_key(ctx *gin.Context) {
	i := ctx.Param("id")
	id, _ := strconv.Atoi(i)
	a.as.Del_key(id)
}
