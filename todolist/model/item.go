package model

import "time"

type ToDoItem struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Title     string     `json:"title" gorm:"column:title;"`
	Status    string     `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (ToDoItem) TableName() string {
	return "todo_items"
}

type DataPaging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}
