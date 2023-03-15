package models

import "time"

type product struct {
  gorm.model
  code  string
  price uint
}
