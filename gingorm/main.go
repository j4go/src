/*
gorm文档：http://gorm.book.jasperxu.com/
*/

package main

import (
	"github.com/gin-gonic/gin"
	"gingorm/handlers"
	//"gingorm/dbops"

)

func main() {
	//defer dbops.DB.Close()

	// route
	r := gin.Default()
	r.GET("/persons", handlers.GetPersons)
	r.GET("/persons/:id", handlers.GetPerson)
	r.POST("/persons", handlers.CreatePerson)
	r.PUT("/persons/:id", handlers.UpdatePerson)
	r.DELETE("/persons/:id", handlers.DeletePerson)

	r.Run(":8080")

}
