package main

import (
	"chilindo/services/user/config"
	"chilindo/services/user/controller"
	"chilindo/services/user/repository"
	"chilindo/services/user/route"
	"chilindo/services/user/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db := config.ConnectDatabase()
	//DI Auth
	r := router()

	userRepo := repository.NewUserRepositoryDefault(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	userRouter := route.NewUserRoute(userController, r)
	userRouter.SetRouter()

	if err := r.Run(":3000"); err != nil {
		log.Println("Open port is fail")
		return
	}
}

func router() *gin.Engine {
	router := gin.Default()
	return router
}
