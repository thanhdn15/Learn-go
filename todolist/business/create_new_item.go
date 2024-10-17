package business

import (
	"context"
	"errors"
	"github.com/thanhdn15/concrete_lean_go/todolist/model"
)

type CreateTodoItemStorage interface {
	CreateItem(ctx context.Context, data *model.ToDoItem) error
}

type createBiz struct {
	store CreateTodoItemStorage
}

func NewCreateTodoItemBiz(store CreateTodoItemStorage) *createBiz {
	return &createBiz{
		store: store,
	}
}

func (biz *createBiz) CreateNewItem(ctx context.Context, data *model.ToDoItem) error {
	if data.Title == "" {
		return errors.New("title cannot be blank")
	}

	data.Status = "Doing"

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}
