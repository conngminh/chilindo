package repository

import (
	"chilindo/src/user-service/entity"
	"gorm.io/gorm"
	"log"
)

type IAddressRepository interface {
	CreateAddress(address *entity.Address) (*entity.Address, error)
}

type AddressRepository struct {
	db *gorm.DB
}

func NewAddressRepositoryDefault(db *gorm.DB) *AddressRepository {
	return &AddressRepository{db: db}
}

func (a *AddressRepository) CreateAddress(address *entity.Address) (*entity.Address, error) {
	var newAddress *entity.Address
	newAddress.UserId = address.UserId
	result := a.db.Create(&newAddress)
	if result.Error != nil {
		log.Println("CreateAddress: Error Create in package repository", result)
		return nil, result.Error
	}
	return newAddress, nil
}
