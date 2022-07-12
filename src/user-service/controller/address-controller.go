package controller

import (
	"chilindo/src/user-service/service"
	"github.com/gin-gonic/gin"
)

type IAddressController interface {
	CreateAddress(c *gin.Context)
	GetAddressByUserId(c *gin.Context)
	UpdateAddress(c *gin.Context)
	DeleteAddress(c *gin.Context)
}

type AddressControllerDefault struct {
	AddressService service.IAddressService
}

func NewAddressControllerDefault(addressService service.IAddressService) *AddressControllerDefault {
	return &AddressControllerDefault{AddressService: addressService}
}

func (a *AddressControllerDefault) CreateAddress(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a *AddressControllerDefault) GetAddressByUserId(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a *AddressControllerDefault) UpdateAddress(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a *AddressControllerDefault) DeleteAddress(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
