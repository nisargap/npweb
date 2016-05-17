// Created By: Nisarga Patel

package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func main() {

	// gets mongo server configuration
	dbConfig := GetDBServerConfig()

	// connects to server
	session, err := mgo.Dial(dbConfig.ServerLocation)

	if err != nil {
		panic(err)
	}
	defer session.Close()

	applicationDB := session.DB(dbConfig.Database)

	if err = applicationDB.Login(dbConfig.Username, dbConfig.Password); err != nil {
		panic(err)
	}

	// Set the session to monotonic behavior
	session.SetMode(mgo.Monotonic, true)

	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", index)
	r.Run() // listen and server on 0.0.0.0:8080
	

}
