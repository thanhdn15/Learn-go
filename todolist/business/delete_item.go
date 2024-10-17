package business

import (
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DeleteItem interface {
	DeleteItemById(ctx context.Context, id int) error
}

type deleteItem struct {
	storage DeleteItem
}

func NewDeleteItem(storage DeleteItem) *deleteItem {
	return &deleteItem{
		storage: storage,
	}
}

func (biz *deleteItem) DeleteItemById(c *gin.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	if err := biz.storage.DeleteItemById(c, id); err != nil {
		return err
	}

	return nil
}
