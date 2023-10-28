package main

import (
	"github.com/crudecample/controllers"
	"github.com/crudecample/inittializers"
	"github.com/gin-gonic/gin"
)

func init() {
	inittializers.ConnectToDB()
	inittializers.LoadEnvvariables()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDeleted)

	r.Run() // listen and serve on 0.0.0.0:3000
}
