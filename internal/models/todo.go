package models

import "gorm.io/gorm"

type Todo struct {
    gorm.Model
    Title  string `gorm:"not null"`
    Status string `gorm:"not null"`
}
