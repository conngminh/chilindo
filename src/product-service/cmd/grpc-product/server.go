package grpc_product

import (
	"chilindo/src/pkg/pb/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

const (
	addr     = ":50052"
	certFile = "src/pkg/ssl/server.crt"
	keyFile  = "src/pkg/ssl/server.pem"
)

type ProductServer struct {
	product.ProductServiceServer
}

func RunGRPCServer(enabledTLS bool, lis net.Listener) error {
	var opts []grpc.ServerOption
	if enabledTLS {
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			return err
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	product.RegisterProductServiceServer(s, &ProductServer{})

	log.Printf("listening on %s\n", addr)
	return s.Serve(lis)
}
