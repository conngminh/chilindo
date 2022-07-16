package main

import (
	grpc_server "chilindo/src/admin-service/cmd/grpc-admin"
	"chilindo/src/admin-service/config"
	"chilindo/src/admin-service/controller"
	"chilindo/src/admin-service/repository"
	"chilindo/src/admin-service/route"
	"chilindo/src/admin-service/service"
	"chilindo/src/admin-service/utils"
	"fmt"
	"log"
	"net"
)

const (
	addr = ":50051"
)

func main() {

	db := config.GetDB()
	newRouter := utils.Router()

	adminRepo := repository.NewAdminRepositoryDefault(db)
	adminService := service.NewAdminServiceDefault(adminRepo)
	adminController := controller.NewAdminControllerDefault(adminService)
	adminRouter := route.NewAdminRouteDefault(adminController, newRouter)
	adminRouter.GetRouter()

	go func() {
		if err := newRouter.Run(":1001"); err != nil {
			fmt.Println("Open port is fail")
			return
		}
		fmt.Println("Server is opened on port 8080")
	}()

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err = grpc_server.RunGRPCServer(true, lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("gRPC server admin is running")

}
