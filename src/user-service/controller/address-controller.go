package controller

import (
	"chilindo/src/user-service/config"
	"chilindo/src/user-service/entity"
	"chilindo/src/user-service/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IAddressController interface {
	CreateAddress(c *gin.Context)
}

type AddressController struct {
	AddressService service.IAddressService
}

func NewAddressControllerDefault(addressService service.IAddressService) *AddressController {
	return &AddressController{AddressService: addressService}
}

func (a *AddressController) CreateAddress(c *gin.Context) {
	var address *entity.Address
	userId, ok := c.Get(config.UserId)
	if !ok {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Error create address",
		})
		log.Println("CreateAddress: Error Get User ID in package controller")
		c.Abort()
		return
	}
	address.UserId = userId.(uint64)
	address, err := a.AddressService.CreateAddress(address)
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"Message": "Error Add address",
		})
		log.Println("CreateAddressBy: Error create new address in package controller")
		return
	}
	c.JSONP(http.StatusOK, address)
}
