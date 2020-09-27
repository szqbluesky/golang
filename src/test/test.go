package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	fmt.Println("hello start to serve!")
	router := gin.Default()
	router.GET("test")
	router.Run("")
}


