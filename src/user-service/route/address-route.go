package route

import (
	"chilindo/src/user-service/controller"
	"github.com/gin-gonic/gin"
)

type IAddressRoute interface {
	GetRouter()
}

type AddressRouteDefault struct {
	AddressController controller.IAddressController
	Route             *gin.Engine
}

func (a *AddressRouteDefault) GetRouter() {
	newAddressRoute(a.AddressController, a.Route)
}

func newAddressRoute(controller controller.IAddressController, group *gin.Engine) {
	addressRoute := group.Group("/chilindo/address")
	{
		addressRoute.POST("/create", controller.CreateAddress)
		addressRoute.GET("/get/:id", controller.GetAddressByUserId)
		addressRoute.PUT("/update", controller.UpdateAddress)
		addressRoute.DELETE("/delete", controller.DeleteAddress)
	}
}
func NewAddressRouteDefault(addressController controller.IAddressController, route *gin.Engine) *AddressRouteDefault {
	return &AddressRouteDefault{AddressController: addressController, Route: route}
}
