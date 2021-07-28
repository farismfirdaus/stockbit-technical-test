package grpc

import (
	"context"
	"database/sql"
	"log"

	"github.com/farismfirdaus/stockbit-technical-test/microservices/logging/model"
	"github.com/farismfirdaus/stockbit-technical-test/microservices/logging/repository"
	logPostgres "github.com/farismfirdaus/stockbit-technical-test/microservices/logging/repository/log"
	pb "github.com/farismfirdaus/stockbit-technical-test/microservices/pb/logging"
)

type logging struct {
	logPg repository.LogInterface
}

func NewLogGrpcHandler(db *sql.DB) pb.LoggingServer {
	return &logging{
		logPg: logPostgres.BuildLog(db),
	}
}

func (l *logging) Log(ctx context.Context, req *pb.Request) (res *pb.Response, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[ERROR] exception: %v", r)
			res.Error = "Internal Server Error"
		}
	}()
	defer func() {
		log.Printf("[ INFO] response: %v", res)
	}()
	log.Printf("[ INFO] request: %v", req)

	err = l.logPg.Insert(&model.DbLog{
		Method:   req.Method,
		Request:  req.Request,
		Response: req.Response,
	})

	if err != nil {
		panic(err.Error())
	}
	return &pb.Response{Error: ""}, nil
}
