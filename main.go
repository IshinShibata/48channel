package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type Todo struct {
	gorm.Model
	Text   string
	Status string
}
type Thread struct {
	gorm.Model
	Content        string
	Title          string
	ThreadContents []ThreadContent
}

type ThreadContent struct {
	gorm.Model
	Content  string
	Title    string
	Thread   Thread
	ThreadID uint
}

func gormConnect() *gorm.DB {
	//user := os.Getenv("MySQL_USER")
	//pass := os.Getenv("MYSQL_PASSWORD")
	//dbname := os.Getenv("MYSQL_DATABASE")
	//host := os.Getenv("MYSQL_HOST")
	//fmt.Println(host)

	//connection := fmt.Sprintf(user, pass, host, dbname, "?charset=utf8&parseTime=True&loc=Local")
	//connection := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)

	//db, err := gorm.Open("mysql", user+":"+
	//	pass+"@tcp(127.0.0.1:3306)/"+dbname+"?charset=utf8&parseTime=True&loc=Local")

	db, err := gorm.Open("mysql", "user:password@tcp(dockerMySQL:3306)/48channel")

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	db := gormConnect()
	defer db.Close()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	db.AutoMigrate(&ThreadContent{}, &Thread{}, &Todo{})

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	router.POST("/threads", insertThread)
	router.GET("/threads", getThreadList)

	router.Run(":3000")

	//todo := Todo{}
	//todo.Text = "test"
	//todo.Status = "testdesu"
	//db.Create(&todo)
}

func insertThread(c *gin.Context) {
	Content := c.PostForm("content")
	Title := c.DefaultPostForm("title", "title")
	db := gormConnect()
	defer db.Close()

	thread := Thread{}
	thread.Content = Content
	thread.Title = Title
	db.Create(&thread)
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"thread": thread,
	})
	// return thread
	// c.JSON(http.StatusOK, gin.H{"firstName": firstName, "lastName": lastName})
}

func getThreadList(c *gin.Context) {
	//firstName := c.PostForm("first_name")
	//lastName := c.DefaultPostForm("last_name", "default_last_name")
	db := gormConnect()
	defer db.Close()

	var todos []Todo
	result := db.Find(&todos)
	fmt.Println(result.Value)
	fmt.Printf("%T\n", result.Value)
	c.JSON(http.StatusOK, result.Value)
}
