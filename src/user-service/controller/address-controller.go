package controller

import (
	"chilindo/src/user-service/config"
	"chilindo/src/user-service/entity"
	"chilindo/src/user-service/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type IAddressController interface {
	CreateAddress(c *gin.Context)
	UpdateAddress(c *gin.Context)
	DeleteAddress(c *gin.Context)
}

type AddressController struct {
	AddressService service.IAddressService
}

func NewAddressControllerDefault(addressService service.IAddressService) *AddressController {
	return &AddressController{AddressService: addressService}
}

func (a *AddressController) CreateAddress(ctx *gin.Context) {
	var newAddress *entity.Address
	errDTO := ctx.ShouldBindJSON(&newAddress)
	if errDTO != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error Binding JSON",
		})
		log.Println("CreateAddress: Error ShouldBindJSON in package controller", errDTO)
		ctx.Abort()
		return
	}

	userId, ok := ctx.Get(config.UserId)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error create address",
		})
		log.Println("CreateAddress: Error Get User ID in package controller")
		ctx.Abort()
		return
	}

	newAddress.UserId = userId.(uint)
	createdAddress, errCreateAddress := a.AddressService.CreateAddress(newAddress)
	if errCreateAddress != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errCreateAddress.Error(),
		})
		log.Println("CreateAddress: Error create new address in package controller")
		return
	}
	ctx.JSON(http.StatusOK, createdAddress)
}

func (a *AddressController) UpdateAddress(c *gin.Context) {
	var updateAddress *entity.Address
	if err := c.ShouldBindJSON(&updateAddress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		log.Println("UpdateAddress: Error ShouldBindJSON in package controller", err)
		return
	}
	userId, ok := c.Get(config.UserId)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error create address",
		})
		log.Println("CreateAddress: Error Get User ID in package controller")
		c.Abort()
		return
	}
	addressId, err := strconv.Atoi(c.Param(config.ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Error create address",
		})
		log.Println("CreateAddress: Error Get User ID in package controller")
		c.Abort()
		return
	}
	updateAddress.ID = uint(addressId)
	updateAddress.UserId = userId.(uint)

	updatedAddress, errUpdate := a.AddressService.UpdateAddress(updateAddress)
	if errUpdate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errUpdate.Error(),
		})
		log.Println("UpdateAddress: Error update address in package controller")
		c.Abort()
		return
	}
	fmt.Println("check here")

	if updatedAddress == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not found address",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, updatedAddress)
}

func (a *AddressController) DeleteAddress(c *gin.Context) {

}