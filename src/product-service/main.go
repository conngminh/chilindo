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
	db                *gorm.DB                     = config.GetDB()
	productRepository repository.ProductRepository = repository.NewProductRepository(db)
	productService    service.ProductService       = service.NewProductService(productRepository)
	productController controller.ProductController = controller.NewProductController(productService)
)

func main() {
	defer config.CloseDatabase(db)
	r := gin.Default()
	//, middleware.AuthorizeJWT(jwtService)
	productRoutes := r.Group("api/product")
	productRoutes.Use(middleware.Logger())
	{
		productRoutes.POST("/create", productController.Insert)
		productRoutes.PUT("/:productId", productController.Update)
		productRoutes.DELETE("/:productId", productController.Delete)
		productRoutes.GET("/:productId", productController.FindByID)
		productRoutes.GET("/", productController.All)
	}
	r.Run()
}
