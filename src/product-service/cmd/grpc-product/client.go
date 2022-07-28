package grpc_product

import (
	"chilindo/src/pkg/pb/admin"
	"chilindo/src/pkg/ssl"
	"google.golang.org/grpc"
	"log"
)

const (
	adminClientPort = "localhost:50051"
)

type IRPCClient interface {
	SetUpProductClient() admin.AdminServiceClient
}

type RPCClient struct{}

func (r RPCClient) SetUpAdminClient() admin.AdminServiceClient {
	var opts []grpc.DialOption
	creds, tlsErr := ssl.LoadTLSCredentials()

	if tlsErr != nil {
		log.Fatalf("Failed to load credentials: %v", tlsErr)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(adminClientPort, opts...)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	log.Println("Listening from port :", adminClientPort)
	adminClient := admin.NewAdminServiceClient(conn)
	return adminClient
}

func NewRPCClient() *RPCClient {
	return &RPCClient{}
}
