package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Id          string `json:"id" validate:"required,min=3,max=10" gorm:"primaryKey"`
	Title       string `json:"title" validate:"required,max=50" gorm:"not null";size:50;"`
	Description string `json:"description"`
	Status      string `json:"status" validate:"required,oneof=todo doing done" gorm:"not null";index;"`
	CreatedAt   int64  `json:"created_at" gorm:"autoCreateTime"`
	DueDate     int64  `json:"due_date"`
	Owner       User   `json:"owner" gorm:"embedded;embeddedPrefix:owner_";`
	Assignee    User   `json:"assignee" gorm:"embedded;embeddedPrefix:assignee_";`
}
