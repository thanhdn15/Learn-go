package storage

import (
	"context"
	"github.com/thanhdn15/concrete_lean_go/todolist/model"
	//"github.com/thanhdn15/concrete_lean_go/todolist/model"
)

func (s *mysqlStorage) GetListOfItems(ctx context.Context, paging model.DataPaging, data *[]model.ToDoItem) (*[]model.ToDoItem, error) {
	if err := s.db.Table(model.ToDoItem{}.TableName()).
		Count(&paging.Total).
		Offset((paging.Page - 1) * paging.Limit).
		Order("id desc").
		Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
