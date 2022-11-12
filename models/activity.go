package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model

	Id    int    `json:"id" gorm:"id"`
	Email string `json:"email" gorm:"email"`
	Title string `json:"title" gorm:"title"`
}
type Activities []Activity

func (Activity) TableName() string {
	return "activity"
}

func (a *Activity) GetAllActivity(DB *gorm.DB) (out *Activities, err error) {

	tx := DB.Table(a.TableName())
	tx.Find(&out)

	if tx.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	if tx.Error != nil {
		return nil, errors.New(tx.Error.Error())
	}

	return
}

func (a *Activity) isActivityExist(DB *gorm.DB) bool {
	var act Activity
	tx := DB.Table(a.TableName())
	tx.First(&act, "id", a.Id)
	if tx.RowsAffected >= 1 {
		return true
	} else {
		return false
	}
}

func (a *Activity) GetOneActivity(DB *gorm.DB) (out *Activity, err error) {

	tx := DB.Table(a.TableName())
	tx.First(&out, "id = ?", a.Id)

	if tx.RowsAffected == 0 {
		return nil, errors.New("data not found")
	}
	if tx.Error != nil {
		return nil, errors.New(tx.Error.Error())
	}

	return
}

func (a *Activity) CreateActivity(DB *gorm.DB) (out *Activity, err error) {

	fmt.Println(a)

	tx := DB.Table(a.TableName()).Create(&Activity{
		Id:    a.Id,
		Email: a.Email,
		Title: a.Title,
	}).Last(&out)

	if tx.Error != nil {
		return nil, errors.New("data not found")
	}

	return
}

func (a *Activity) UpdateActivity(DB *gorm.DB) (out *Activity, err error) {

	tx := DB.Table(a.TableName()).Updates(&Activity{
		Id:    a.Id,
		Email: a.Email,
		Title: a.Title,
	}).Last(&out)

	if tx.Error != nil {
		return nil, errors.New("data not found")
	}

	return
}

func (a *Activity) DeleteActivity(DB *gorm.DB) (err error) {

	tx := DB.Table(a.TableName()).Delete(&a)
	if tx.RowsAffected == 0 {
		return errors.New("data not found")
	}
	if tx.Error != nil {
		return errors.New("delete failed")
	}
	return
}
