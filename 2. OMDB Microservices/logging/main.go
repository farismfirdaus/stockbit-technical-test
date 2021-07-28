package main

import (
	"database/sql"
	"log"
	"net"

	logGrpc "github.com/farismfirdaus/stockbit-technical-test/microservices/logging/delivery/grpc"
	pb "github.com/farismfirdaus/stockbit-technical-test/microservices/pb/logging"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	log.Printf("open database: %v", viper.GetString("database.url"))
	db, err := sql.Open("postgres", viper.GetString("database.url"))
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Failed to listen on port 9001: %v", err)
	}

	s := logGrpc.NewLogGrpcHandler(db)
	grpcServer := grpc.NewServer()
	pb.RegisterLoggingServer(grpcServer, s)

	log.Println("serving 9001")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over post 9001: %v", err)
	}
}
