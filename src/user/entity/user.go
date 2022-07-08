package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname string `json:"firstname" gorm:"type:nvarchar(100);not null"`
	Lastname  string `json:"description" gorm:"type:nvarchar(100);not null"`
	Username  string `json:"username" gorm:"type:nvarchar(100);not null"`
	Password  string `json:"password" gorm:"type:nvarchar(100);not null"`
	Birthday  string `json:"birthday" gorm:"type:nvarchar(100)"`
	Phone     string `json:"phone" gorm:"type:nvarchar(100)"`
	Email     string `json:"email" gorm:"type:nvarchar(100)"`
	Gender    bool   `json:"gender" gorm:"type:boolean"`
	Country   string `json:"country" gorm:"type:nvarchar(100)"`
	Language  string `json:"language" gorm:"type:nvarchar(100)"`
}
