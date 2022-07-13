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
