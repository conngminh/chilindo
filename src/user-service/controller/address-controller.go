package controller

import (
	"chilindo/src/user-service/config"
	"chilindo/src/user-service/entity"
	"chilindo/src/user-service/service"
	"fmt"
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
	var newAddress *entity.Address
	userId, ok := c.Get(config.UserId)
	fmt.Println("neeeeeee", userId)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error create address",
		})
		log.Println("CreateAddress: Error Get User ID in package controller")
		c.Abort()
		return
	}
	newAddress.UserId = userId.(uint)
	createdAddress, err := a.AddressService.CreateAddress(newAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error Add address",
		})
		log.Println("CreateAddress: Error create new address in package controller")
		return
	}
	c.JSON(http.StatusOK, createdAddress)

}
