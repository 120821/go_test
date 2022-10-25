package models

import (
	"time"
)

type Apple struct {
	Id        int       `json:"id" gorm:"column:id;unique;primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"column:name"`
  Color     int       `json:"color" gorm:"column:color"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

//定义表名
func (Users) TableName() string {
	return "apples"
}

