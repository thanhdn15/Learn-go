package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhdn15/concrete_lean_go/todolist/business"
	storage2 "github.com/thanhdn15/concrete_lean_go/todolist/storage"
	"gorm.io/gorm"
	"net/http"
)

func HandleEditItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		storage := storage2.NewMysqlStorage(db)
		biz := business.NewEditItem(storage)

		data, err := biz.EditItemData(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
