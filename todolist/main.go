package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"
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
		v1.POST("/items", createItem(db))
		v1.GET("/items", getListOfItems(db))
		v1.GET("/items/:id", readItemById(db))
		v1.POST("/items/:id", editItemById(db))
		v1.DELETE("/items/:id", deleteItemById(db))
	}

	router.Run()
}

func createItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem ToDoItem

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		dataItem.Title = strings.TrimSpace(dataItem.Title)

		if dataItem.Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "title cannot be blank"})

			return
		}

		dataItem.Status = "Doing"

		if err := db.Create(&dataItem).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"data": dataItem.Id,
		})
	}
}

func getListOfItems(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		type DataPaging struct {
			Page  int   `json:"page" form:"page"`
			Limit int   `json:"limit" form:"limit"`
			Total int64 `json:"total" form:"-"`
		}

		var paging DataPaging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if paging.Page <= 0 {
			paging.Page = 1
		}

		if paging.Limit <= 0 {
			paging.Limit = 10
		}

		offset := (paging.Page - 1) * paging.Limit

		var result []ToDoItem

		if err := db.Table(ToDoItem{}.TableName()).
			Count(&paging.Total).
			Offset(offset).
			Order("id desc").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}

func readItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem ToDoItem

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Where("id = ?", id).First(&dataItem).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": dataItem,
		})
	}
}

func editItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		var dataItem ToDoItem

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Where("id = ?", id).Updates(&dataItem).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}

func deleteItemById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Table(ToDoItem{}.TableName()).
			Where("id = ?", id).
			Delete(nil).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}
