package service

import (
	"chilindo/src/user-service/entity"
	"chilindo/src/user-service/repository"
	"log"
)

type IAddressService interface {
	CreateAddress(address *entity.Address) (*entity.Address, error)
	UpdateAddress(address *entity.Address) (*entity.Address, error)
	DeleteAddress(id string) error
}

type AddressService struct {
	AddressRepository repository.IAddressRepository
}

func NewAddressServiceDefault(addressRepository repository.IAddressRepository) *AddressService {
	return &AddressService{AddressRepository: addressRepository}
}

func (a *AddressService) CreateAddress(address *entity.Address) (*entity.Address, error) {
	newAddress, err := a.AddressRepository.CreateAddress(address)
	if err != nil {
		log.Println("CreateAddress: Error Create address in package service", err)
		return nil, err
	}
	return newAddress, nil
}

func (a *AddressService) UpdateAddress(address *entity.Address) (*entity.Address, error) {
	updateAddress, err := a.AddressRepository.UpdateAddress(address)
	if err != nil {
		log.Println("UpdateAddress: Error Update address in package service", err)
		return nil, err
	}
	return updateAddress, nil
}

func (a *AddressService) DeleteAddress(id string) error {
	err := a.AddressRepository.DeleteAddress(id)
	if err != nil {
		log.Println("DeleteAddress: Error delete address in package service", err)
		return err
	}
	return nil
}
