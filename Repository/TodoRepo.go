package Repository

import (
	"github.com/souvik666/go_todo/Config"
	"github.com/souvik666/go_todo/Models"
)

func CreateTodo(todo *Models.Todo) error {
	var err error = Config.DB.Create(todo).Error
	if err != nil {
		return err
	}
	return nil
}

func GetALlTodo(todo *[]Models.Todo) error {
	err := Config.DB.Find(todo).Error
	if err != nil {
		return err
	}
	return nil
}

func GetOneTodo(todo *Models.Todo, id string) error {
	err := Config.DB.Where("id= ?", id).First(todo).Error
	if err != nil {
		return err
	}
	return nil

}
func UpdateOneTodo(todo *Models.Todo) error {
	err := Config.DB.Save(todo).Error
	if err != nil {
		return err
	}
	return nil

}

func DeleteOneTodo(todo *Models.Todo, id string) error {
	err := Config.DB.Where("id= ?", id).Delete(todo).Error
	if err != nil {
		return err
	}
	return nil

}
