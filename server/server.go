package server

import (
	todo "github.com/EikoNakashima/test-go2/controller"
	"github.com/gin-gonic/gin"
)

// Init is initialize server
func Init() {
	r := router()
	r.LoadHTMLGlob("views/*.html")
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("views/*.html")

	u := r.Group("/todos")
	{
		ctrl := todo.Controller{}
		u.GET("", ctrl.Index)
		u.GET("/:id", ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	return r
}
