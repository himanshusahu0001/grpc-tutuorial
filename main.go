package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/akuity/grpc-gateway-client/pkg/grpc/gateway"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/user/test/proto"
	"google.golang.org/grpc"
)

type MyServiceServer struct {
	pb.UnimplementedMyServiceServer
}

func (s MyServiceServer) GetMessage(ctx context.Context, req *pb.RequestStruct) (*pb.ResponseStruct, error) {
	return &pb.ResponseStruct{Text: "Hello form " + req.ClientType + " client"}, nil
}

func StartGrpcServer(myServiceServer pb.MyServiceServer) string {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterMyServiceServer(grpcServer, myServiceServer)

	serverAddress := listener.Addr().String()

	go grpcServer.Serve(listener)

	log.Println("GRPC server started at", serverAddress)
	return serverAddress
}

func StartRestServer(myServiceServer pb.MyServiceServer) string {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
	}
	mux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterMyServiceHandlerServer(ctx, mux, myServiceServer)
	if err != nil {
		fmt.Println(err)
	}

	serverAddress := listener.Addr().String()

	go http.Serve(listener, mux)

	log.Println("REST server started at", serverAddress)
	return serverAddress
}

func main() {
	myServiceServer := MyServiceServer{}

	grpcServerAddress := StartGrpcServer(myServiceServer)

	restServerAddress := StartRestServer(myServiceServer)

	// gRPC client code -------------------------
	conn, err := grpc.Dial(grpcServerAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	grpcClient := pb.NewMyServiceClient(conn)
	req := &pb.RequestStruct{ClientType: "gRPC"}
	res, err := grpcClient.GetMessage(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Text)

	// REST Client------------------------
	restServerAddress = "http://0.0.0.0:8081" // grpc-gateway server address
	restClient := pb.NewMyServiceGatewayClient(gateway.NewClient(restServerAddress))
	req.ClientType = "REST"
	res, err = restClient.GetMessage(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Text)

	// ------------------------
	var toKeepProgramRunning string
	fmt.Scanln(&toKeepProgramRunning)
}
