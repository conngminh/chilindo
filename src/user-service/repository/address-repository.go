package repository

import (
	"chilindo/src/user-service/entity"
	"gorm.io/gorm"
	"log"
)

type AddressRepository interface {
	CreateAddress(address *entity.Address) (*entity.Address, error)
	GetAddressByUserID(userID uint) (*[]entity.Address, error)
	UpdateAddress(address *entity.Address) (*entity.Address, error)
	DeleteAddress(address *entity.Address) error
}

type AddressRepositoryDefault struct {
	db *gorm.DB
}

func NewAddressRepositoryDefault(db *gorm.DB) *AddressRepositoryDefault {
	return &AddressRepositoryDefault{db: db}
}

func (a *AddressRepositoryDefault) CreateAddress(address *entity.Address) (*entity.Address, error) {
	record := a.db.Create(&address)
	if record.Error != nil {
		log.Println("CreateAddress: Create address", record.Error)
		return address, record.Error
	}
	return address, nil
}

//
func (a *AddressRepositoryDefault) GetAddressByUserID(userID uint) (*[]entity.Address, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AddressRepositoryDefault) UpdateAddress(address *entity.Address) (*entity.Address, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AddressRepositoryDefault) DeleteAddress(address *entity.Address) error {
	//TODO implement me
	panic("implement me")
}
