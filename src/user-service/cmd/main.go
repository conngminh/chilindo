package main

import (
	"chilindo/src/user-service/config"
	"chilindo/src/user-service/controller"
	"chilindo/src/user-service/repository"
	"chilindo/src/user-service/route"
	"chilindo/src/user-service/service"
	"chilindo/src/user-service/utils"
	"log"
	"net/http"
)

func main() {
	db := config.GetDB()
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
	http.Handle("/", newRouter)
	log.Fatal(http.ListenAndServe(":3000", newRouter))
}
