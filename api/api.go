package api

import (
	"app/core"
	"app/usecase"
	"net/http"
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
	id, err := strconv.Atoi(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	s := a.as.Post(id, name)
	ctx.JSON(http.StatusOK, gin.H{
		"message": s,
	})
}

func (a *Api) Get_key(ctx *gin.Context) {
	i := ctx.Param("key")
	id, err := strconv.Atoi(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid key",
		})
		return
	}

	id1, name := a.as.Get_key(id)
	ctx.JSON(http.StatusOK, gin.H{
		"id":   id1,
		"name": name,
	})
}

func (a *Api) Get_Print(ctx *gin.Context) {
	var output []core.User
	var st string
	output, st = a.as.Get_Print()

	if st == "nil" {
		ctx.JSON(http.StatusOK, gin.H{
			"print": output,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"error": st,
		})
	}

}

func (a *Api) Del_all(ctx *gin.Context) {
	output := a.as.Del_all()
	ctx.JSON(http.StatusOK, gin.H{
		"output": output,
	})
}

func (a *Api) Del_key(ctx *gin.Context) {
	i := ctx.Param("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	output := a.as.Del_key(id)
	ctx.JSON(http.StatusOK, gin.H{
		"output": output,
	})
}
