package pb

//go:generate protoc  service.proto -I ../third_party/ -I ./      --go_out=. --go_opt=paths=source_relative       --go-grpc_out=. --go-grpc_opt=paths=source_relative     --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative     --openapiv2_out=./swagger
