package main

import (
	"fmt"

	"github.com/mrasif/websocketdemo/socketservice"

	"github.com/gin-gonic/gin"
)

func main() {
	go socketservice.GetHub().Run()

	router := gin.New()
	router.LoadHTMLFiles("index.html")

	router.GET("/room/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/ws/:roomId", func(c *gin.Context) {
		fmt.Println(c.GetQuery("token"))
		roomId := c.Param("roomId")
		if roomId != "1" {
			c.JSON(200, gin.H{
				"message": "No room found",
			})
			return
		}
		socketservice.ServeWs(c.Writer, c.Request, roomId)
	})

	router.Run("0.0.0.0:8080")
}
