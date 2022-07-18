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
	optionRoutes := o.Router.Group("api/product/option")
	optionRoutes.Use(middleware.Logger())
	{
		optionRoutes.POST("/:productId", o.ProductOptionController.CreateOption)
		optionRoutes.GET("/:productId", o.ProductOptionController.GetOptions)
		optionRoutes.DELETE("/:optionId", o.ProductOptionController.DeleteOption)
		optionRoutes.GET("/option/:optionId", o.ProductOptionController.GetOptionByID)
		optionRoutes.PUT("/:optionId", o.ProductOptionController.UpdateOption)

	}
}
