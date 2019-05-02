package service

import (
	"github.com/gin-gonic/gin"
	//	"github.com/go-xorm/xorm"
)

type UserService struct {
}

func (service *UserService) Login(ctx *gin.Context, username string, pwd string) (err error) {

	return
}
