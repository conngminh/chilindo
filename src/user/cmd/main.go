package main

import (
	"chilindo/src/user/config"
	"chilindo/src/user/controller"
	"chilindo/src/user/repository"
	"chilindo/src/user/route"
	"chilindo/src/user/service"
	"chilindo/src/user/utils"
	"log"
)

func main() {
	db := config.ConnectDatabase()
	//DI Auth
	newRouter := utils.Router()

	userRepo := repository.NewUserRepositoryDefault(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	userRouter := route.NewUserRoute(userController, newRouter)
	userRouter.SetRouter()

	if err := newRouter.Run(":3000"); err != nil {
		log.Println("Open port is fail")
		return
	}
}
