package business

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/thanhdn15/concrete_lean_go/todolist/model"
)

type GetListItemsStorage interface {
	GetListOfItems(ctx context.Context, paging model.DataPaging, data *[]model.ToDoItem) (*[]model.ToDoItem, error)
}

type listItemsStorage struct {
	store GetListItemsStorage
}

func NewGetListItemStorage(store GetListItemsStorage) *listItemsStorage {
	return &listItemsStorage{
		store: store,
	}
}

func (biz *listItemsStorage) GetListItems(c *gin.Context) (*[]model.ToDoItem, error) {
	var paging model.DataPaging
	var data *[]model.ToDoItem

	if err := c.ShouldBind(&paging); err != nil {
		return nil, err
	}

	if paging.Page <= 0 {
		paging.Page = 1
	}

	if paging.Limit <= 0 {
		paging.Limit = 10
	}

	data, err := biz.store.GetListOfItems(c, paging, data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
