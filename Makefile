run:
	go run cmd/main.go

# This command generates the code from .proto files in proto directory and swagger documentation in docs directory
pb:
	protoc --go_out=./ --proto_path=./proto ./proto/search.proto
	protoc --go-grpc_out=./ --proto_path=./proto ./proto/search.proto
	protoc --grpc-gateway_out=./ --openapiv2_out ./docs --proto_path=./proto ./proto/search.proto

# This command installs dependencies for working with grpc
install:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/protobuf/cmd/protoc-gen-go