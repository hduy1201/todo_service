package models

type User struct {
	Id        string `json:"id" validate:"required,min=3,max=10" gorm:"primaryKey"`
	Username  string `json:"username" validate:"required,max=50" gorm:"not null";size:50;uniqueIndex;"`
	HPassword string `json:"hpassword" validate:"required" gorm:"not null"`
}
