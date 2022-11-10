package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	Id              int            `json:"id" gorm:"id"`
	ActivityGroupId string         `json:"activity_group_id" gorm:"activity_group_id"`
	Title           string         `json:"title" gorm:"title"`
	IsActive        string         `json:"is_active" gorm:"is_active"`
	Priority        string         `json:"priority" gorm:"priority"`
	CreatedAt       time.Time      `json:"created_at" gorm:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at" gorm:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"deleted_at"`
}
type Todos []Todo

func (Todo) TableName() string {
	return "todo"
}

func (t *Todo) GetAllTodo(DB *gorm.DB) (out *Todos, err error) {

	tx := DB.Table(t.TableName())
	tx.Find(&out)

	if tx.Error != nil {
		return nil, errors.New("data not found")
	}

	return
}

func (t *Todo) GetOneTodo(DB *gorm.DB) (out *Todos, err error) {

	tx := DB.Table(t.TableName())
	tx.First(&out, "id = ?", t.Id)

	if tx.Error != nil {
		return nil, errors.New("data not found")
	}

	return
}

func (t *Todo) CreateTodo(DB *gorm.DB) (out *Todos, err error) {

	tx := DB.Table(t.TableName()).Create(&Todo{
		Id:              t.Id,
		ActivityGroupId: t.ActivityGroupId,
		Title:           t.Title,
		IsActive:        t.IsActive,
		Priority:        t.Priority,
		CreatedAt:       t.CreatedAt,
		UpdatedAt:       t.UpdatedAt,
		DeletedAt:       t.DeletedAt,
	}).Last(&out)

	if tx.Error != nil {
		return nil, errors.New("data not found")
	}

	return
}

func (t *Todo) UpdateTodo(DB *gorm.DB) (out *Todos, err error) {

	tx := DB.Table(t.TableName()).Updates(&Todo{
		Id:              t.Id,
		ActivityGroupId: t.ActivityGroupId,
		Title:           t.Title,
		IsActive:        t.IsActive,
		Priority:        t.Priority,
		CreatedAt:       t.CreatedAt,
		UpdatedAt:       t.UpdatedAt,
		DeletedAt:       t.DeletedAt,
	}).Last(&out)

	if tx.Error != nil {
		return nil, errors.New("data not found")
	}

	return
}

func (t *Todo) DeleteTodo(DB *gorm.DB) (err error) {

	tx := DB.Table(t.TableName()).Delete(&t)

	if tx.Error != nil {
		return errors.New("data not found")
	}
	return
}
