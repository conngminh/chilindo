package route

import (
	"chilindo/src/product-service/controller"
	"chilindo/src/product-service/middleware"
	"github.com/gin-gonic/gin"
)

type IProductOptionRoute interface {
	GetRouter()
}

type OptionRoute struct {
	ProductOptionController controller.ProductOptionController
	Router                  *gin.Engine
}

func NewOptionRoute(productOptionController controller.ProductOptionController, router *gin.Engine) *OptionRoute {
	return &OptionRoute{ProductOptionController: productOptionController, Router: router}
}

func (o OptionRoute) GetRouter() {
	optionRoutes := o.Router.Group("api/option")
	optionRoutes.Use(middleware.Logger())
	{
		optionRoutes.POST("/:productId/product", o.ProductOptionController.CreateOption)
		optionRoutes.GET("/:productId/product", o.ProductOptionController.GetOptions)
		optionRoutes.DELETE("/:optionId", o.ProductOptionController.DeleteOption)
		//optionRoutes.GET("/:optionId", o.ProductOptionController.GetOptionByID)
		optionRoutes.PUT("/:optionId", o.ProductOptionController.UpdateOption)

	}
}
