package entity

type Administrator struct {
	Id       uint   `json:"id" gorm:"primary_key;auto_increment"`
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Email    string `json:"email" gorm:"type:nvarchar(100);not null"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Token    string `gorm:"-" json:"token,omitempty"`
}
