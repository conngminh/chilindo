package grpc_admin

import (
	"chilindo/src/admin-service/config"
	"chilindo/src/admin-service/repository"
	"chilindo/src/admin-service/service"
	"chilindo/src/pkg/pb/admin"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

const (
	addrAdminServer = ":50051"
	certFile        = "src/pkg/ssl/server.crt"
	keyFile         = "src/pkg/ssl/server.pem"
)

type AdminServer struct {
	admin.AdminServiceServer
	AdminService service.IAdminService
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

	adminRepo := repository.NewAdminRepositoryDefault(config.DB)
	AdminService := service.NewAdminServiceDefault(adminRepo)

	admin.RegisterAdminServiceServer(s, &AdminServer{
		AdminService: AdminService,
	})

	log.Printf("Admin Server is on port %s\n", addrAdminServer)
	return s.Serve(lis)
}

func (a *AdminServer) CheckIsAuth(ctx context.Context, in *admin.CheckIsAuthRequest) (*admin.CheckIsAuthResponse, error) {
	log.Printf("Login request: %v\n", in)

	res, err := a.AdminService.CheckIsAuth(in)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %v", err)
	}

	if res == nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return res, nil
}
