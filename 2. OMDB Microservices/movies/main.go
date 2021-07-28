package main

import (
	"log"
	"net"

	moviesGrpc "github.com/farismfirdaus/stockbit-technical-test/microservices/movies/delivery/grpc"
	"github.com/farismfirdaus/stockbit-technical-test/microservices/pb/logging"
	pb "github.com/farismfirdaus/stockbit-technical-test/microservices/pb/movies"
	"google.golang.org/grpc"
)

func main() {
	conn2, err := grpc.Dial("localhost:9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn2.Close()
	logClient := logging.NewLoggingClient(conn2)

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := moviesGrpc.NewMoviesGrpcHandler(logClient)
	grpcServer := grpc.NewServer()
	pb.RegisterMoviesServer(grpcServer, s)

	log.Println("serving 9000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over post 9000: %v", err)
	}
}
