package repository

import (
	"chilindo/src/user-service/entity"
	"gorm.io/gorm"
	"log"
)

type IAddressRepository interface {
	CreateAddress(address *entity.Address) (*entity.Address, error)
}

type AddressRepositoryDefault struct {
	db *gorm.DB
}

func NewAddressRepositoryDefault(db *gorm.DB) *AddressRepositoryDefault {
	return &AddressRepositoryDefault{db: db}
}

func (a AddressRepositoryDefault) CreateAddress(address *entity.Address) (*entity.Address, error) {
	if errCheckEmptyField := address.Validate(); errCheckEmptyField != nil {
		log.Println("CreateAddress: Error empty field in package repository", errCheckEmptyField)
		return nil, errCheckEmptyField
	}
	result := a.db.Create(&address)
	if result.Error != nil {
		log.Println("CreateAddress: Error Create in package repository", result)
		return nil, result.Error
	}
	return address, nil

}
