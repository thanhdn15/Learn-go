package storage

import (
	"context"
	"github.com/thanhdn15/concrete_lean_go/todolist/model"
)

func (s *mysqlStorage) CreateItem(ctx context.Context, data *model.ToDoItem) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
