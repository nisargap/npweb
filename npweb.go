package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", index)
	r.Run() // listen and server on 0.0.0.0:8080

}
