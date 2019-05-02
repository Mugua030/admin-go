package router

import (
	"admin-go/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("views/**/*")
	//r.Static("/assets", "./assets")
	r.StaticFS("/assets", http.Dir("./assets"))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404 , page not exists!",
		})
	})
	//r.Use(ValidateToken())
	//Login AND logout
	auth := new(controller.AuthController)
	r.GET("/login", auth.ToLogin)
	r.POST("/login", auth.Login)

	//Index
	index := new(controller.IndexController)
	r.GET("/", index.Home)

	/*
		admin := r.Group("/admin")
		{
			admin.GET("/register", handlers.Register)
			admin.GET("/index", handlers.Index)
			admin.GET("/login", handlers.LoginGet)
			admin.POST("/login", handlers.LoginPost)
				admin.GET("/login", func(c *gin.Context) {
					c.HTML(http.StatusOK, "admin/login.tmpl", gin.H{
						"title": "mmm88",
					})
				})
		}
	*/

	return r
}

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.FormValue("token")
		if token == "" {
			c.JSON(401, gin.H{
				"msg": "toen required",
			})
			c.Abort()
			return
		}
		if token != "accesstoken" {
			c.JSON(http.StatusOK, gin.H{
				"msg": "invalid token",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
