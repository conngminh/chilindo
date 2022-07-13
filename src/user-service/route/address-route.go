package route

import (
	"chilindo/src/user-service/controller"

	"chilindo/src/user-service/middleware"
	"github.com/gin-gonic/gin"
)

type IAddressRoute interface {
	GetRouter()
}

type AddressRoute struct {
	AddressController controller.IAddressController
	Router            *gin.Engine
	JWTMiddleware     *middleware.SMiddleWare
}

func NewAddressRouteDefault(addressController controller.IAddressController, router *gin.Engine) *AddressRoute {
	return &AddressRoute{AddressController: addressController, Router: router}
}

func (a AddressRoute) GetRouter() {
	addressRoute := a.Router.Group("/chilindo/user/address").Use(a.JWTMiddleware.IsAuthenticated())
	{
		addressRoute.POST("/create", a.AddressController.CreateAddress)
	}
}
