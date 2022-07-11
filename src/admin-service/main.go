package main

import (
	"chilindo/src/admin-service/config"
	"chilindo/src/admin-service/controller"
	"chilindo/src/admin-service/repository"
	"chilindo/src/admin-service/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB                   = config.SetupDatabaseConnection()
	adminRepository repository.AdminRepository = repository.NewAdminRepository(db)
	jwtService      service.JWTService         = service.NewJWTService()
	adminService    service.AdminService       = service.NewAdminService(adminRepository)
	adminController controller.AdminController = controller.NewAdminController(adminService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	//, middleware.AuthorizeJWT(jwtService)
	adminRoutes := r.Group("api/admin")
	{
		adminRoutes.POST("/login", adminController.Login)
		adminRoutes.POST("/register", adminController.Register)
	}
}
