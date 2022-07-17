package main

import (
	"chilindo/src/auction-service/controller"

	"chilindo/src/auction-service/config"
	//admin_server_controller "chilindo/src/product-service/controller/admin-grpc-controller"
	"chilindo/src/auction-service/repository"
	"chilindo/src/auction-service/route"
	"chilindo/src/auction-service/service"
	"chilindo/src/user-service/utils"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

const (
	//DB_CONNECTION_STRING = "DB_CONNECTION_STRING"
	ginPort        = ":1009"
	grpcClientPort = "localhost:50051"
	grpcServerPort = "localhost:50052"
	certFile       = "src/pkg/ssl/ca.crt"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		return nil, err
	}
	return creds, nil
}

func main() {

	var opts []grpc.DialOption
	creds, err := loadTLSCredentials()

	if err != nil {
		log.Fatalf("Failed to load credentials: %v", err)
	}

	opts = append(opts, grpc.WithTransportCredentials(creds))

	//grpc CLient
	conn, err := grpc.Dial(grpcClientPort, opts...)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	//adminClient := admin.NewAdminServiceClient(conn)

	//Product service DB
	db := config.GetDB()
	defer config.CloseDatabase(db)
	newRouter := utils.Router()

	auctionRepository := repository.NewAuctionRepositoryDefault(db)
	auctionService := service.NewAuctionServiceDefault(auctionRepository)
	auctionController := controller.NewAuctionController(auctionService)
	//adminSrvCtrl := admin_server_controller.NewAdminServiceController()
	//auctionRouter := route.NewProductRoute(productController, newRouter, adminSrvCtrl, adminClient)
	auctionRouter := route.NewAuctionRoute(auctionController, newRouter)
	auctionRouter.GetRouter()

	if err := newRouter.Run(ginPort); err != nil {
		fmt.Println("Open port is fail")
		return
	}
	fmt.Println("Server is opened on port 8080")

	//go func() {
	//	if err := newRouter.Run(ginPort); err != nil {
	//
	//		fmt.Println("Open port is fail")
	//		return
	//	}
	//	fmt.Println("Server is opened on port 8080")
	//}()
	//
	//lis, err := net.Listen("tcp", grpcServerPort)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//
	//if err = grpc_product.RunGRPCServer(true, lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
	//log.Println("gRPC server admin is running")
}
