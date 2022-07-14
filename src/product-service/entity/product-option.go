package entity

type ProductOption struct {
	Id           string  `json:"id" gorm:"primary_key;auto_increment"`
	ProductId    string  `json:"productId" gorm:"type:varchar(20);not null"`
	Color        string  `json:"color" gorm:"type:nvarchar(100)"`
	Size         string  `json:"size" gorm:"type:nvarchar(100)"`
	Model        string  `json:"model" gorm:"type:nvarchar(100)"`
	ProductModel string  `json:"productModel"`
	Product      Product `json:"-" gorm:"foreignKey:ProductId;references:Id"`
}
