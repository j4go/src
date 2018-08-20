package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Person struct {
	Name     string    `form:"name" json:"name" xml:"name" binding:"required"`
	Address  string    `form:"address" json:"address" xml:"address" binding:"required"`
	Birthday time.Time `form:"birthday" json:"birthday" josn:"birthday" time_format:"2006-01-02" binding:"required"`
}

func main() {
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()



	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name+" is "+action)
	})

	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Fan")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	r.POST("/form_post", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Fan")
		lastname := c.Query("lastname")
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "Fan")
		c.JSON(http.StatusOK, gin.H{
			"status":   "posted",
			"message":  message,
			"nick":     nick,
			"firtname": firstname,
			"lastname": lastname,
		})
	})

	/*
		curl -X POST http://localhost:8080/upload \
		-F "file=@/Users/molock/gostudy.zip" \
		-H "Content-Type: multipart/form-data"
	*/
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("%s uploaded!\n", file.Filename))
	})

	r.Any("/testing", startPage)

	r.GET("/json", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"name"`
			Message string `json:"msg"`
			Number  int    `json:"number"`
		}
		msg.Name = "lgf"
		msg.Message = "I'm message"
		msg.Number = 999
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/jsonarray", func(c *gin.Context) {
		names := []string{"aaa", "bbb", "ccc"}
		c.SecureJSON(http.StatusOK, names)
	})

	r.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			//"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})


	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "This is my title",
		})
	})


	r.Run()
}

func startPage(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err == nil {
		log.Println("====== Bind Info ======")
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)

	} else {
		log.Println(err)
	}
	c.String(200, "Success")
}
