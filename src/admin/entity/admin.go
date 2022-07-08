package entity

type Administrator struct {
	Id       uint   `json:"id" gorm:"primary_key;auto_increment"`
	Username string `json:"username" gorm:"type:nvarchar(100);not null"`
	Password string `json:"password"gorm:"type:nvarchar(100);not null"`
}
