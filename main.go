package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/mrasif/websocketdemo/socketservice"

	"github.com/gin-gonic/gin"
)

func main() {
	go socketservice.GetHub().Run()

	router := gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = []string{"Accept", "Accept-CH", "Accept-Charset", "Accept-Datetime", "Accept-Encoding", "Accept-Ext", "Accept-Features", "Accept-Language", "Accept-Params", "Accept-Ranges", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Origin", "Access-Control-Expose-Headers", "Access-Control-Max-Age", "Access-Co  ntrol-Request-Headers", "Access-Control-Request-Method", "Authorization", "Content-Type"}
	corsConfig.AllowAllOrigins = true

	// setup cors middleware
	router.Use(cors.New(corsConfig))

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
