package main

import (
	"chilindo/src/pkg/pb/admin"
	grpc_product "chilindo/src/product-service/cmd/grpc-product"
	"chilindo/src/product-service/config"
	"chilindo/src/product-service/controller"
	admin_server_controller "chilindo/src/product-service/controller/admin-grpc-controller"
	"chilindo/src/product-service/repository"
	"chilindo/src/product-service/route"
	"chilindo/src/product-service/service"
	"chilindo/src/user-service/utils"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

const (
	//DB_CONNECTION_STRING = "DB_CONNECTION_STRING"
	ginPort        = ":1002"
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

	adminClient := admin.NewAdminServiceClient(conn)

	//Product service DB
	db := config.GetDB()
	defer config.CloseDatabase(db)
	newRouter := utils.Router()

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)
	adminSrvCtrl := admin_server_controller.NewAdminServiceController()
	productRouter := route.NewProductRoute(productController, newRouter, adminSrvCtrl, adminClient)
	productRouter.GetRouter()

	productOptionRepository := repository.NewProductOptionRepository(db)
	productOptionService := service.NewProductOptionService(productOptionRepository)
	productOptionController := controller.NewProductOptionController(productOptionService)
	optionRouter := route.NewOptionRoute(productOptionController, newRouter)
	optionRouter.GetRouter()

	productImageRepository := repository.NewProductImageRepository(db)
	productImageService := service.NewProductImageService(productImageRepository)
	productImageController := controller.NewProductImageController(productImageService)
	imageRouter := route.NewImageRoute(productImageController, newRouter)
	imageRouter.GetRouter()

	//if err := newRouter.Run(ginPort); err != nil {
	//
	//	fmt.Println("Open port is fail")
	//	return
	//}
	//fmt.Println("Server is opened on port 8080")

	go func() {
		if err := newRouter.Run(ginPort); err != nil {

			fmt.Println("Open port is fail")
			return
		}
		fmt.Println("Server is opened on port 8080")
	}()

	lis, err := net.Listen("tcp", grpcServerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err = grpc_product.RunGRPCServer(true, lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("gRPC server admin is running")
}
