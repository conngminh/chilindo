package entity

type Product struct {
	Id          uint    `json:"id" gorm:"primary_key;auto_increment"`
	Name        string  `json:"name" gorm:"type:nvarchar(100);not null"`
	Min_price   float64 `json:"min_price" gorm:"type:double;not null"`
	Description string  `json:"description" gorm:"type:nvarchar(500);not null"`
	Quantity    int     `json:"quantity" gorm:"type:nvarchar(500);not null"`
}

type ProductOption struct {
	Id          uint    `json:"id" gorm:"primary_key;auto_increment"`
	Name        string  `json:"name" gorm:"type:nvarchar(100);not null"`
	Min_price   float64 `json:"min_price" gorm:"type:double;not null"`
	Description string  `json:"description" gorm:"type:nvarchar(500);not null"`
	Quantity    int     `json:"quantity" gorm:"type:nvarchar(500);not null"`
}
