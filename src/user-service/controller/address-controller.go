package controller

import (
	"chilindo/src/admin-service/helper"
	"chilindo/src/user-service/config"
	"chilindo/src/user-service/entity"
	"chilindo/src/user-service/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IAddressController interface {
	CreateAddress(c *gin.Context)
	GetAddressByUserId(c *gin.Context)
	UpdateAddress(c *gin.Context)
	DeleteAddress(c *gin.Context)
}

type AddressControllerDefault struct {
	AddressService service.IAddressService
	jwtService     service.JWTService
}

func NewAddressControllerDefault(addressService service.IAddressService, jwtService service.JWTService) *AddressControllerDefault {
	return &AddressControllerDefault{AddressService: addressService, jwtService: jwtService}
}

func (a *AddressControllerDefault) CreateAddress(c *gin.Context) {
	var newAddress *entity.Address
	err := c.ShouldBind(&newAddress)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, res)
	} else {
		//authHeader := c.GetHeader("Authorization")
		//tokenSigned := strings.TrimPrefix(authHeader, "bearer ")
		userId, oke := c.Get(config.UserId)
		log.Println(userId)
		if !oke {
			c.JSONP(http.StatusBadRequest, gin.H{
				"Message": "Error create address",
			})
			log.Println("CreateAddressByUserId: Error Get User ID in package controller")
			c.Abort()
			return
		}
		newAddress.UserID = userId.(uint64)
		result, _ := a.AddressService.CreateAddress(newAddress)
		response := helper.BuildResponse(true, "OK", result)
		c.JSON(http.StatusCreated, response)
	}
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
