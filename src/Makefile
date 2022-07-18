run-adminsrv:
	go run src/admin-service/cmd/main.go

run-productsrv:
	go run src/product-service/cmd/main.go

gen-protoc-admin:
	protoc --go_out=. --go-grpc_out=. .\src\pkg\proto\admin.proto

gen-protoc-product:
	protoc --go_out=. --go-grpc_out=. .\src\pkg\proto\product.proto

<<<<<<< HEAD
=======
mock-gen-controller:
	mockgen -source="./src/product-service/repository/product-repository.go" -destination="./src/product-service/mock/product-repository_mock.go" -package=service

>>>>>>> 5dfe3c89866a7c6b428c3658a6190f718dead0dd
