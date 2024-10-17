package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhdn15/concrete_lean_go/todolist/business"
	"github.com/thanhdn15/concrete_lean_go/todolist/model"
	storage2 "github.com/thanhdn15/concrete_lean_go/todolist/storage"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func HandleCreateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem model.ToDoItem

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		dataItem.Title = strings.TrimSpace(dataItem.Title)

		storage := storage2.NewMysqlStorage(db)
		biz := business.NewCreateTodoItemBiz(storage)
		
		if err := biz.CreateNewItem(c.Request.Context(), &dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": dataItem.Id,
		})
	}
}
