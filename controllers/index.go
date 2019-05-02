package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
	name string
}

func (IndexController) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index.tmpl", gin.H{
		"title": "mmmuuu",
	})
}
