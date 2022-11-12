package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model

	Id              int            `json:"id" gorm:"id"`
	ActivityGroupId int            `json:"activity_group_id" gorm:"activity_group_id"`
	Title           string         `json:"title" gorm:"title"`
	IsActive        *bool          `json:"is_active" gorm:"is_active"`
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
	tx.First(&out, "id = ?", t.Id)

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
		return nil, errors.New("create failed")
	}

	return
}

func (t *Todo) UpdateTodo(DB *gorm.DB) (out *Todos, err error) {

	fmt.Println(t)

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

	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}
	if tx.Error != nil {
		return errors.New("delete failed")
	}
	return
}
