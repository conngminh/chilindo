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

//package main
//
//import (
//rpcClient "chilindo/src/auction-service/cmd/grpc-auction"
//"chilindo/src/auction-service/config"
//"chilindo/src/auction-service/controller"
//"chilindo/src/auction-service/repository"
//"chilindo/src/auction-service/route"
//"chilindo/src/auction-service/service"
//"chilindo/src/pkg/pb/admin"
//admin_server_controller "chilindo/src/product-service/controller/admin-grpc-controller"
//"chilindo/src/user-service/utils"
//"fmt"
//"google.golang.org/grpc"
//"google.golang.org/grpc/credentials"
//"log"
//)
//
//const (
//	ginPort        = ":1009"
//	grpcClientPort = "localhost:50053"
//	grpcServerPort = "localhost:50052"
//	certFile       = "src/pkg/ssl/ca.crt"
//)
//
//func loadTLSCredentials() (credentials.TransportCredentials, error) {
//	creds, err := credentials.NewClientTLSFromFile(certFile, "")
//	if err != nil {
//		return nil, err
//	}
//	return creds, nil
//}
//
//func main() {
//
//	var opts []grpc.DialOption
//	creds, err := loadTLSCredentials()
//
//	if err != nil {
//		log.Fatalf("Failed to load credentials: %v", err)
//	}
//
//	opts = append(opts, grpc.WithTransportCredentials(creds))
//
//	//grpc CLient
//	conn, err := grpc.Dial(grpcClientPort, opts...)
//	if err != nil {
//		log.Fatalf("failed to connect: %v", err)
//	}
//	defer conn.Close()
//
//	adminClient := admin.NewAdminServiceClient(conn)
//	//Create new gRPC Client
//	grpcClient := rpcClient.NewRPCClient()
//	productClient := grpcClient.SetUpProductClient()
//
//	//Product service DB
//	db := config.GetDB()
//	defer config.CloseDatabase(db)
//	newRouter := utils.Router()
//
//	auctionRepository := repository.NewAuctionRepositoryDefault(db)
//	auctionService := service.NewAuctionServiceDefault(auctionRepository)
//	auctionController := controller.NewAuctionController(auctionService, productClient)
//	adminSrvCtrl := admin_server_controller.NewAdminServiceController()
//	auctionRouter := route.NewAuctionRoute(auctionController, newRouter, adminSrvCtrl, adminClient)
//	auctionRouter.GetRouter()
//
//	if err := newRouter.Run(ginPort); err != nil {
//		fmt.Println("Open port is fail")
//		return
//	}
//	fmt.Println("Server is opened on port 1009")
//}
