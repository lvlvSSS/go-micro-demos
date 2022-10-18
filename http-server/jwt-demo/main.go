package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/jwt", func(c *gin.Context) {
		data, _ := c.GetRawData()
		log.Printf("%s \n", string(data))
		c.Writer.WriteHeader(200)
		c.Writer.Header().Set("Content-Type", "text")
		c.Writer.Write([]byte("success"))
	})
	r.POST("/login", func(c *gin.Context) {
		token := jwt.New(jwt.SigningMethodHS256)
		signedString, err := token.SignedString([]byte("123456"))
		if err != nil {
			c.JSON(404, gin.H{
				"message": "not conrrect",
			})
		}
		c.Writer.WriteHeader(200)
		c.Writer.Header().Set("token", signedString)
		c.Writer.Header().Set("Content-Type", "text")
		c.Writer.Write([]byte("success"))
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
