package entity

type Product struct {
	Id          string  `json:"id" gorm:"primary_key;type:varchar(20);not null"`
	Name        string  `json:"name" gorm:"type:nvarchar(100);not null"`
	MinPrice    float64 `json:"minPrice" gorm:"type:double;not null"`
	Description string  `json:"description" gorm:"type:nvarchar(500);not null"`
	Quantity    int     `json:"quantity" gorm:"type:nvarchar(500);not null"`
}
