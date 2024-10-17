package storage

import (
	"context"
	"github.com/thanhdn15/concrete_lean_go/todolist/model"
)

func (s *mysqlStorage) DeleteItemById(ctx context.Context, id int) error {
	if err := s.db.Table(model.ToDoItem{}.TableName()).
		Where("id = ?", id).
		Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
