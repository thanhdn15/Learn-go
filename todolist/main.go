package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhdn15/concrete_lean_go/todolist/transport"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type ToDoItem struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Title     string     `json:"title" gorm:"column:title;"`
	Status    string     `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (ToDoItem) TableName() string {
	return "todo_items"
}

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/todo_items?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySql: ", err)
	}

	log.Println("Connected", db)

	router := gin.Default()

	v1 := router.Group("v1")
	{
		v1.POST("/items", transport.HandleCreateItem(db))
		v1.GET("/items", transport.HandleGetListOfItems(db))
		v1.GET("/items/:id", transport.HandleReadByItemById(db))
		v1.POST("/items/:id", transport.HandleEditItem(db))
		v1.DELETE("/items/:id", transport.HandleDeleteItem(db))
	}

	router.Run()
}
