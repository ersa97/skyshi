package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	Id        int            `json:"id" gorm:"id"`
	Email     string         `json:"email" gorm:"email"`
	Title     string         `json:"title" gorm:"title"`
	CreatedAt time.Time      `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"deleted_at"`
}
type Activities []Activity

func (Activity) TableName() string {
	return "activity"
}

func (a *Activity) GetAllActivity(DB *gorm.DB) (out *Activities, err error) {

	tx := DB.Table(a.TableName())
	tx.Find(&out)

	if tx.Error != nil {
		return nil, errors.New("data not found")
	}

	return
}

func (a *Activity) GetOneActivity(DB *gorm.DB) (out *Activity, err error) {

	tx := DB.Table(a.TableName())
	tx.First(&out, "id = ?", a.Id)

	if tx.Error != nil {
		return nil, errors.New("data not found")
	}

	return
}

func (a *Activity) CreateActivity(DB *gorm.DB) (out *Activity, err error) {

	fmt.Println(a)

	tx := DB.Table(a.TableName()).Create(&Activity{
		Id:        a.Id,
		Email:     a.Email,
		Title:     a.Title,
		CreatedAt: a.CreatedAt,
		DeletedAt: a.DeletedAt,
	}).Last(&out)

	if tx.Error != nil {
		return nil, errors.New("data not found")
	}

	return
}

func (a *Activity) UpdateActivity(DB *gorm.DB) (out *Activity, err error) {

	tx := DB.Table(a.TableName()).Updates(&Activity{
		Id:        a.Id,
		Email:     a.Email,
		Title:     a.Title,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		DeletedAt: a.DeletedAt,
	}).Last(&out)

	if tx.Error != nil {
		return nil, errors.New("data not found")
	}

	return
}

func (a *Activity) DeleteActivity(DB *gorm.DB) (err error) {

	tx := DB.Table(a.TableName()).Delete(&a)

	if tx.Error != nil {
		return errors.New("data not found")
	}
	return
}
