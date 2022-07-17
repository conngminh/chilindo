package main

import (
	rpcClient "chilindo/src/auction-service/cmd/grpc-auction"
	"chilindo/src/auction-service/config"
	"chilindo/src/auction-service/controller"
	"chilindo/src/auction-service/repository"
	"chilindo/src/auction-service/route"
	"chilindo/src/auction-service/service"
	"chilindo/src/user-service/utils"
	"fmt"
)

const (
	ginPort = ":1009"
)

func main() {
	//Create new gRPC Client
	grpcClient := rpcClient.NewRPCClient()
	productClient := grpcClient.SetUpProductClient()

	//Product service DB
	db := config.GetDB()
	defer config.CloseDatabase(db)
	newRouter := utils.Router()

	auctionRepository := repository.NewAuctionRepositoryDefault(db)
	auctionService := service.NewAuctionServiceDefault(auctionRepository)
	auctionController := controller.NewAuctionController(auctionService, productClient)
	auctionRouter := route.NewAuctionRoute(auctionController, newRouter)
	auctionRouter.GetRouter()

	if err := newRouter.Run(ginPort); err != nil {
		fmt.Println("Open port is fail")
		return
	}
	fmt.Println("Server is opened on port 8080")
}
