package Models

import (
	"github.com/souvik666/go_todo/Constants"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Status      Constants.TodoStatus `json:"status"`
}

func (b *Todo) TableName() Constants.Tables {
	return "todo"
}
