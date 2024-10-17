package business

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/thanhdn15/concrete_lean_go/todolist/model"
	"strconv"
)

type ReadByItemId interface {
	ReadByItem(ctx context.Context, id int, data *model.ToDoItem) (*model.ToDoItem, error)
}

type readByItemId struct {
	storage ReadByItemId
}

func NewReadByItemId(storage ReadByItemId) *readByItemId {
	return &readByItemId{
		storage: storage,
	}
}

func (biz *readByItemId) ReadItemById(c *gin.Context) (*model.ToDoItem, error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return nil, err
	}

	var data *model.ToDoItem
	data, err = biz.storage.ReadByItem(c, id, data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
