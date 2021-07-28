package main

import (
	"log"

	deliveryHttp "github.com/farismfirdaus/stockbit-technical-test/microservices/api-gateway/delivery/http"
	logging "github.com/farismfirdaus/stockbit-technical-test/microservices/pb/logging"
	movie "github.com/farismfirdaus/stockbit-technical-test/microservices/pb/movies"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	movieClient := movie.NewMoviesClient(conn)

	conn2, err := grpc.Dial("localhost:9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn2.Close()
	logClient := logging.NewLoggingClient(conn2)

	r := gin.Default()
	deliveryHttp.NewMoviesHttpHandler(r, movieClient, logClient)

	if err := r.Run(":8052"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
