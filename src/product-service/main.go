package main

import (
	"chilindo/src/product-service/config"
	"chilindo/src/product-service/controller"
	"chilindo/src/product-service/middleware"
	"chilindo/src/product-service/repository"
	"chilindo/src/product-service/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                      *gorm.DB                           = config.GetDB()
	productRepository       repository.ProductRepository       = repository.NewProductRepository(db)
	productImageRepository  repository.ProductImageRepository  = repository.NewProductImageRepository(db)
	productOptionRepository repository.ProductOptionRepository = repository.NewProductOptionRepository(db)
	productService          service.ProductService             = service.NewProductService(productRepository)
	productImageService     service.ProductImageService        = service.NewProductImageService(productImageRepository)
	productOptionService    service.ProductOptionService       = service.NewProductOptionService(productOptionRepository)
	productController       controller.ProductController       = controller.NewProductController(productService)
	productImageController  controller.ProductImageController  = controller.NewProductImageController(productImageService)
	productOptionController controller.ProductOptionController = controller.NewProductOptionController(productOptionService)
)

func main() {
	defer config.CloseDatabase(db)
	r := gin.Default()
	productRoutes := r.Group("api/product")
	productRoutes.Use(middleware.Logger())
	{
		productRoutes.POST("/create", productController.Insert)
		productRoutes.PUT("/:productId", productController.Update)
		productRoutes.DELETE("/:productId", productController.Delete)
		productRoutes.GET("/:productId", productController.FindByID)
		productRoutes.GET("/", productController.All)
		productRoutes.POST("/:productId/option", productOptionController.CreateOption)
		productRoutes.GET("/:productId/option", productOptionController.GetOptions)
		productRoutes.POST("/:productId/image", productImageController.CreateImage)
		productRoutes.GET("/:productId/image", productImageController.GetImage)

	}
	optionRoutes := r.Group("api/option")
	optionRoutes.Use(middleware.Logger())
	{
		optionRoutes.DELETE("/:optionId", productOptionController.DeleteOption)
		optionRoutes.GET("/:optionId", productOptionController.GetOptionByID)
		optionRoutes.PUT("/:optionId", productOptionController.UpdateOption)
	}
	imageRoutes := r.Group("api/image")
	imageRoutes.Use(middleware.Logger())
	{
		imageRoutes.DELETE("/:imageId", productImageController.DeleteImage)
		imageRoutes.GET("/:imageId", productImageController.GetImageByID)
		imageRoutes.PUT("/:imageId", productImageController.UpdateImage)
	}
	r.Run()
}
