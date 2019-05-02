package controller

import (
	"admin-go/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
}

var daoServiceUser service.UserService

type User struct {
	Username string `json:username`
	Password string `json:password`
}

func (ac *AuthController) ToLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/login.tmpl", gin.H{
		"title": "login in",
	})
}
func (ac *AuthController) Login(ctx *gin.Context) {

	var u User
	ctx.BindJSON(&u)
	if u.Username != "mugua" {
		ctx.JSON(http.StatusOK, gin.H{
			"errno":  "2",
			"errmsg": fmt.Sprintf("fail data: %s", u.Username),
		})
		return
	}
	username := u.Username
	password := u.Password

	err := daoServiceUser.Login(ctx, username, password)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"errno":  "2",
			"errmsg": fmt.Sprintf("fail data: %s", u.Username),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"errno":  "0",
		"errmsg": "success",
	})
}
