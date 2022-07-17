package ssl

import "google.golang.org/grpc/credentials"

const certFile = "src/pkg/ssl/ca.crt"

func LoadTLSCredentials() (credentials.TransportCredentials, error) {
	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		return nil, err
	}
	return creds, nil
}
