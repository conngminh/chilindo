package entity

import (
	"gorm.io/gorm"
)

type Auction struct {
	gorm.Model `json:"-"`
	//Id         int     `json:"id" gorm:"primary_key;type:nvarchar(20);not null"`
	StartTime  string  `json:"start-time" gorm:"type:nvarchar(100);not null"`
	EndTime    string  `json:"end-ime" gorm:"type:nvarchar(100);not null"`
	CurrentBid float64 `json:"current-bid" gorm:"type:nvarchar(100);not null"`
	Quantity   int     `json:"quantity" gorm:"type:nvarchar(100);not null"`
}
