package entity

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Firstname   string `json:"firstname" gorm:"type:nvarchar(100);not null"`
	Lastname    string `json:"description" gorm:"type:nvarchar(100);not null"`
	Phone       string `json:"phone" gorm:"type:nvarchar(100)"`
	Email       string `json:"email" gorm:"type:nvarchar(100); not null"`
	Province    string `json:"province" gorm:"type:nvarchar(100); not null"`
	District    string `json:"district" gorm:"type:nvarchar(100); not null"`
	SubDistrict string `json:"sub_district" gorm:"type:nvarchar(100)"`
	Address     string `json:"address" gorm:"type:nvarchar(200); not null"`
	TypeAddress string `gorm:"type_address" gorm:"type:nvarchar(100); not null"`
	UserID      uint64 `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"json:"-"`
}
