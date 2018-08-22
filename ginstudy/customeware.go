package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)


// Custom Middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")
		//log.Println("Print something.")
		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main(){
	r := gin.New()
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println(example)
		c.String(200, "Done")
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}