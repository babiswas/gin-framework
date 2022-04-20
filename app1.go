package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
}

func main() {
	fmt.Println("Server Started")
	users := []User{User{ID: 2, Name: "Tapan", Title: "Biswas"}, User{ID: 2, Name: "Tapan", Title: "Biswas"}}
	fmt.Println(users)
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/example", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/user.html", gin.H{
			"data": users,
		})
	})
	router.Run(":8900")
}
