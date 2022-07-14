package entity

type ProductImages struct {
	Id        uint    `json:"id" gorm:"primary_key"`
	ProductId string  `json:"productId" gorm:"type:varchar(20);not null"`
	Link      string  `json:"link" gorm:"type:varchar(100)"`
	Product   Product `json:"-" gorm:"foreignKey:ProductId"`
}
