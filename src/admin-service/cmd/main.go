package main

import (
	"chilindo/src/admin-service/config"
	"chilindo/src/admin-service/controller"
	"chilindo/src/admin-service/repository"
	"chilindo/src/admin-service/route"
	"chilindo/src/admin-service/service"
	"chilindo/src/admin-service/utils"
	"fmt"
)

func main() {
	db := config.GetDB()
	newRouter := utils.Router()

	adminRepo := repository.NewAdminRepositoryDefault(db)
	adminService := service.NewAdminServiceDefault(adminRepo)
	adminController := controller.NewAdminControllerDefault(adminService)
	adminRouter := route.NewAdminRouteDefault(adminController, newRouter)
	adminRouter.GetRouter()

	if err := newRouter.Run(":8080"); err != nil {

		fmt.Println("Open port is fail")
		return
	}
	fmt.Println("Server is opened on port 8080")
}
