package todo

import (
	"strconv"

	"github.com/EikoNakashima/test-go2/db"
	"github.com/gin-gonic/gin"
)

// Controller is user controlller
type Controller struct{}

//Index
func (pc Controller) Index(c *gin.Context) {
	// router.GET("/", func(ctx *gin.Context) {

	todos := db.GetAll()
	c.HTML(200, "index.html", gin.H{
		"todos": todos,
	})
	// })
}

//Create
func (pc Controller) Create(c *gin.Context) {
	// router.POST("/new", func(ctx *gin.Context) {
	text := c.PostForm("text")
	status := c.PostForm("status")
	db.Insert(text, status)
	c.Redirect(302, "/")
	// })
}

//Detail
func (pc Controller) Show(c *gin.Context) {
	// router.GET("/detail/:id", func(ctx *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	todo := db.GetOne(id)
	c.HTML(200, "detail.html", gin.H{"todo": todo})
	// })
}

//Update
func (pc Controller) Update(c *gin.Context) {
	// router.POST("/update/:id", func(ctx *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	text := c.PostForm("text")
	status := c.PostForm("status")
	db.Update(id, text, status)
	c.Redirect(302, "/")
	// })
}

//削除確認
func (pc Controller) ShowDelete(c *gin.Context) {
	// router.GET("/delete_check/:id", func(ctx *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	todo := db.GetOne(id)
	c.HTML(200, "delete.html", gin.H{"todo": todo})
	// })
}

//Delete
func (pc Controller) Delete(c *gin.Context) {
	// router.POST("/delete/:id", func(ctx *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("ERROR")
	}
	db.Delete(id)
	c.Redirect(302, "/")

	// })
}
