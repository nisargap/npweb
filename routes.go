package main

import "github.com/gin-gonic/gin"
import "net/http"

func index(c *gin.Context) {

	/*********************************
	  Example of sending JSON data
	  	c.JSON(200, gin.H{
	  		"message": "hello",
	  	})
	  *********************************/
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "Website Title",
	})

}
