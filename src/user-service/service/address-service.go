package service

import (
	"chilindo/src/user-service/entity"
	"chilindo/src/user-service/repository"
)

type IAddressService interface {
	CreateAddress(address entity.Address) (entity.Address, error)
	GetAddressByUserID(userID uint) ([]entity.Address, error)
	UpdateAddress(address entity.Address) (entity.Address, error)
	DeleteAddress(address entity.Address) error
}

type AddressServiceDefault struct {
	AddressRepository repository.AddressRepository
}

func NewAddressServiceDefault(addressRepository repository.AddressRepository) *AddressServiceDefault {
	return &AddressServiceDefault{AddressRepository: addressRepository}
}

func (ad *AddressServiceDefault) CreateAddress(address entity.Address) (entity.Address, error) {
	//TODO implement me
	panic("implement me")
}

func (ad *AddressServiceDefault) GetAddressByUserID(userID uint) ([]entity.Address, error) {
	//TODO implement me
	panic("implement me")
}

func (ad *AddressServiceDefault) UpdateAddress(address entity.Address) (entity.Address, error) {
	//TODO implement me
	panic("implement me")
}

func (ad *AddressServiceDefault) DeleteAddress(address entity.Address) error {
	//TODO implement me
	panic("implement me")
}
