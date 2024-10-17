package storage

import (
	"context"
	"github.com/thanhdn15/concrete_lean_go/todolist/model"
)

func (s *mysqlStorage) ReadByItem(ctx context.Context, id int, data *model.ToDoItem) (*model.ToDoItem, error) {
	if err := s.db.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
