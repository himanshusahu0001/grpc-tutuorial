package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/user/test/proto"
	"google.golang.org/grpc"
)

type MyServiceServer struct {
	pb.UnimplementedMyServiceServer
}

func (s MyServiceServer) GetMessage(ctx context.Context, req *pb.RequestStruct) (*pb.ResponseStruct, error) {
	return &pb.ResponseStruct{Text: "gRPC API Working"}, nil
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

func StartRestServer(myServiceServer pb.MyServiceServer) error {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
	}
	mux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterMyServiceHandlerServer(ctx, mux, myServiceServer)
	if err != nil {
		return err
	}

	log.Println("REST server started at", listener.Addr().String())
	return http.Serve(listener, mux)
}

func main() {
	myServiceServer := MyServiceServer{}

	serverAddress := StartGrpcServer(myServiceServer)

	err := StartRestServer(myServiceServer)
	if err != nil {
		fmt.Println(err)
	}

	// client code
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	client := pb.NewMyServiceClient(conn)
	req := &pb.RequestStruct{}
	res, err := client.GetMessage(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Text)

	var toKeepProgramRunning string
	fmt.Scanln(&toKeepProgramRunning)
}
