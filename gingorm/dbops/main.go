package dbops


import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gingorm/models"
)


var DB *gorm.DB
var err error


func init() {
	DB, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gingorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//defer DB.Close()  不能用在这里

	DB.SingularTable(true) //全局设置表名不可以为复数形式

	DB.SetMaxIdleConns(20)
	DB.SetMaxOpenConns(200)

	// init db with db model
	DB.AutoMigrate(&models.Person{})
}
