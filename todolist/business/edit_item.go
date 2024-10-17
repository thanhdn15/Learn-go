package business

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/thanhdn15/concrete_lean_go/todolist/model"
	"strconv"
)

type EditItem interface {
	UpdateItem(ctx context.Context, id int, data *model.ToDoItem) (*model.ToDoItem, error)
}

type editItem struct {
	storage EditItem
}

func NewEditItem(storage EditItem) *editItem {
	return &editItem{
		storage: storage,
	}
}

func (biz *editItem) EditItemData(c *gin.Context) (*model.ToDoItem, error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return nil, err
	}

	var dataItem *model.ToDoItem

	if err := c.ShouldBind(&dataItem); err != nil {
		return nil, err
	}

	data, err := biz.storage.UpdateItem(c, id, dataItem)

	if err != nil {
		return nil, err
	}

	return data, nil
}
