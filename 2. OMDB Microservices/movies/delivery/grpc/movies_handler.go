package grpc

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/farismfirdaus/stockbit-technical-test/microservices/movies/model"
	"github.com/farismfirdaus/stockbit-technical-test/microservices/pb/logging"
	pb "github.com/farismfirdaus/stockbit-technical-test/microservices/pb/movies"
)

type server struct {
	logClient logging.LoggingClient
}

func NewMoviesGrpcHandler(logClient logging.LoggingClient) pb.MoviesServer {
	return &server{
		logClient: logClient,
	}
}

func (s *server) Find(ctx context.Context, req *pb.Request) (res *pb.Response, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[ERROR] exception: %v", r)
			res.Response = "False"
			res.Error = "Internal Server Error"
		}
	}()
	defer func() {
		log.Printf("[ INFO] response: %v", res)
	}()
	log.Printf("[ INFO] request: %v", req)

	params := url.Values{}
	params.Add("s", req.Keywords)
	params.Add("page", req.Pagination)

	data := hitAPI(params)

	request, _ := json.Marshal(req)
	response, _ := json.Marshal(data)
	s.logClient.Log(ctx, &logging.Request{
		Method:   "GRPC",
		Request:  string(request),
		Response: string(response),
	})

	return data.ToMovies(), nil
}

func (s *server) FindByImdbID(ctx context.Context, req *pb.Request) (res *pb.ResponseImdbID, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[ERROR] exception: %v", r)
			res.Response = "False"
			res.Error = "Internal Server Error"
		}
	}()
	defer func() {
		log.Printf("[ INFO] response: %v", res)
	}()
	log.Printf("[ INFO] request: %v", req)

	params := url.Values{}
	params.Add("i", req.ImdbID)

	data := hitAPI(params)

	request, _ := json.Marshal(req)
	response, _ := json.Marshal(data)
	s.logClient.Log(ctx, &logging.Request{
		Method:   "GRPC",
		Request:  string(request),
		Response: string(response),
	})

	return data.ToDetail(), nil
}

func hitAPI(params url.Values) model.Response {
	params.Add("apikey", "faf7e5bb")
	buffer := bytes.NewBufferString("http://www.omdbapi.com/")
	buffer.WriteString("?" + params.Encode())

	httpReq, err := http.NewRequest("GET", buffer.String(), nil)
	if err != nil {
		panic("create request: ")
	}

	httpClient := http.Client{Timeout: time.Second * 10}
	resp, err := httpClient.Do(httpReq)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	response := map[string]interface{}{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response
}
