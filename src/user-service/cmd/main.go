package main

import (
	"chilindo/src/user-service/config"
	"chilindo/src/user-service/controller"
	"chilindo/src/user-service/repository"
	"chilindo/src/user-service/route"
	"chilindo/src/user-service/service"
	"chilindo/src/user-service/utils"
	"fmt"
)

func main() {
	db := config.GetDB()
	newRouter := utils.Router()

	userRepo := repository.NewUserRepositoryDefault(db)
	userService := service.NewUserServiceDefault(userRepo)
	userController := controller.NewUserControllerDefault(userService, service.NewJWTService())
	userRouter := route.NewUserRouteDefault(userController, newRouter)
	userRouter.GetRouter()

	if err := newRouter.Run(":3000"); err != nil {
		fmt.Println("Open port is fail")
		return
	}
	fmt.Println("Server is opened on port 8080")
}
