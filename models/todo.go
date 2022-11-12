package models

import (
	"errors"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model

	ActivityGroupId int    `json:"activity_group_id" gorm:"activity_group_id"`
	Title           string `json:"title" gorm:"title"`
	IsActive        *bool  `json:"is_active" gorm:"is_active"`
	Priority        string `json:"priority" gorm:"priority"`
}
type Todos []Todo

func (Todo) TableName() string {
	return "todo"
}

func (t *Todo) GetAllTodo(DB *gorm.DB) (out *Todos, err error) {

	tx := DB.Table(t.TableName())
	tx.Find(&out, "activity_group_id", t.ActivityGroupId)

	if tx.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	if tx.Error != nil {
		return nil, errors.New(tx.Error.Error())
	}

	return
}

func (t *Todo) GetOneTodo(DB *gorm.DB) (out *Todos, err error) {

	tx := DB.Table(t.TableName())
	tx.First(&out, "id = ?", t.ID)

	if tx.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	if tx.Error != nil {
		return nil, errors.New(tx.Error.Error())
	}

	return
}

func (t *Todo) CreateTodo(DB *gorm.DB) (out *Todos, err error) {

	var activity Activity

	activity.Id = t.ActivityGroupId

	isExist := activity.isActivityExist(DB)
	if !isExist {
		return nil, errors.New("activity group doesn't exist")
	}

	tx := DB.Table(t.TableName()).Create(&Todo{
		ActivityGroupId: t.ActivityGroupId,
		Title:           t.Title,
		IsActive:        t.IsActive,
		Priority:        t.Priority,
	}).Last(&out)

	if tx.Error != nil {
		return nil, errors.New("create failed")
	}

	return
}

func (t *Todo) UpdateTodo(DB *gorm.DB) (out *Todos, err error) {

	tx := DB.Table(t.TableName()).Updates(&Todo{
		ActivityGroupId: t.ActivityGroupId,
		Title:           t.Title,
		IsActive:        t.IsActive,
		Priority:        t.Priority,
	}).Last(&out)

	if tx.Error != nil {
		return nil, errors.New("data not found")
	}

	return
}

func (t *Todo) DeleteTodo(DB *gorm.DB) (err error) {

	tx := DB.Table(t.TableName()).Delete(&t)

	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}
	if tx.Error != nil {
		return errors.New("delete failed")
	}
	return
}
