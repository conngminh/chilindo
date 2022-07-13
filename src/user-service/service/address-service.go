package service

import (
	"chilindo/src/user-service/entity"
	"chilindo/src/user-service/repository"
	"log"
)

type IAddressService interface {
	CreateAddress(address *entity.Address) (*entity.Address, error)
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
