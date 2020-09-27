package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	// 初始化gin
	router := gin.Default()
	//router.Static("/static", "static")
	//router.LoadHTMLFiles("tmplates/*")
	router.GET("/test", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message":"test"})
	})
	router.GET("/getparam", getParam)
	router.POST("/post", post)

	err := router.Run(":80")
	if nil != err{
		println(err)
	}
}

func post(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message":"post"})
}

func getParam(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message":"getParam"})
}