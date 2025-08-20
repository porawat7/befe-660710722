package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

func main(){
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		user :=[]User{{ID:"1", Name:"Nuttachot"}}
		c.JSON(200, user)
	})

	r.Run(":8080")
}
