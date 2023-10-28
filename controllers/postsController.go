package controllers

import (
	"github.com/crudecample/inittializers"
	model "github.com/crudecample/module"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	//get data from request body
	var body struct {
		Body  string `json:body`
		Title string `json:title`
	}

	c.Bind(&body)

	//create post

	//return it
	post := model.Post{Title: body.Title, Body: body.Body}
	result := inittializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//get the posts
	var posts []model.Post
	inittializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostShow(c *gin.Context) {
	id := c.Param("id")
	var posts []model.Post
	inittializers.DB.First(&posts, id)
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostUpdate(c *gin.Context) {
	//get the id off the url
	id := c.Param("id")

	//get the data of req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//find the post were updating
	var posts model.Post
	inittializers.DB.First(&posts, id)

	//update it
	inittializers.DB.Model(&posts).Updates(model.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostDeleted(c *gin.Context) {
	//get the id off the url
	id := c.Param("id")

	inittializers.DB.Delete(&model.Post{}, id)
	//response
	c.Status(200)
}
