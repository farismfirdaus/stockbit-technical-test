package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/farismfirdaus/stockbit-technical-test/microservices/pb/logging"
	"github.com/farismfirdaus/stockbit-technical-test/microservices/pb/movies"
	"github.com/gin-gonic/gin"
)

func NewMoviesHttpHandler(r *gin.Engine, client movies.MoviesClient, logClient logging.LoggingClient) {
	r.GET("/api/movies", getMovies(logClient))
	r.GET("/api/movie/detail", getMovieDetail(logClient))
	r.GET("/api/grpc/movies", grpcMovies(client, logClient))
	r.GET("/api/grpc/movie/detail", grpcMovieDetail(client, logClient))
}

func getMovies(logClient logging.LoggingClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqParams := c.Request.URL.Query()
		params := url.Values{}
		if keyword, ok := reqParams["searchword"]; ok && keyword != nil {
			for _, v := range keyword {
				params.Add("s", v)
			}
		}
		if pagination, ok := reqParams["pagination"]; ok && pagination != nil {
			for _, v := range pagination {
				params.Add("page", v)
			}
		}

		response := hitAPI(params)
		c.JSON(http.StatusOK, response)

		req, _ := json.Marshal(reqParams)
		res, _ := json.Marshal(response)
		logClient.Log(c, &logging.Request{
			Method:   "REST",
			Request:  string(req),
			Response: string(res),
		})
	}
}

func getMovieDetail(logClient logging.LoggingClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqParams := c.Request.URL.Query()
		params := url.Values{}
		if imdbId, ok := reqParams["imdbId"]; ok && imdbId != nil {
			for _, v := range imdbId {
				params.Add("i", v)
			}
		}
		response := hitAPI(params)
		c.JSON(http.StatusOK, response)

		req, _ := json.Marshal(reqParams)
		res, _ := json.Marshal(response)
		logClient.Log(c, &logging.Request{
			Method:   "REST",
			Request:  string(req),
			Response: string(res),
		})
	}
}

func grpcMovies(client movies.MoviesClient, logClient logging.LoggingClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqMovies := movies.Request{}
		reqParams := c.Request.URL.Query()
		if keyword, ok := reqParams["searchword"]; ok && keyword != nil {
			reqMovies.Keywords = keyword[0]
		}
		if pagination, ok := reqParams["pagination"]; ok && pagination != nil {
			reqMovies.Pagination = pagination[0]
		}

		data, err := client.Find(c, &reqMovies)
		if err != nil {
			panic(err.Error())
		}
		c.JSON(http.StatusOK, data)

		req, _ := json.Marshal(reqParams)
		res, _ := json.Marshal(data)
		logClient.Log(c, &logging.Request{
			Method:   "REST-GRPC",
			Request:  string(req),
			Response: string(res),
		})
	}
}

func grpcMovieDetail(client movies.MoviesClient, logClient logging.LoggingClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqDetail := movies.Request{}
		reqParams := c.Request.URL.Query()
		if imdbId, ok := reqParams["imdbId"]; ok && imdbId != nil {
			reqDetail.ImdbID = imdbId[0]
		}

		data, err := client.FindByImdbID(c, &reqDetail)
		if err != nil {
			panic(err.Error())
		}
		c.JSON(http.StatusOK, data)

		req, _ := json.Marshal(reqParams)
		res, _ := json.Marshal(data)
		logClient.Log(c, &logging.Request{
			Method:   "REST-GRPC",
			Request:  string(req),
			Response: string(res),
		})
	}
}

func hitAPI(params url.Values) map[string]interface{} {
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
