package model

import (
	"strings"

	pb "github.com/farismfirdaus/stockbit-technical-test/microservices/pb/movies"
)

type ResponseMovies struct {
	Search       []*Movie `json:"Search"`
	TotalResults string   `json:"totalResults"`
	Response     string   `json:"Response"`
	Error        string   `json:"Error"`
}

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type Response map[string]interface{}

func (r Response) ToMovies() *pb.Response {
	if value, ok := r["Response"].(string); value != "" && ok {
		if strings.EqualFold(value, "true") {
			movies := []*pb.Search{}
			search := r["Search"].([]interface{})
			for _, v := range search {
				container := v.(map[string]interface{})
				movies = append(movies, &pb.Search{
					Title:  container["Title"].(string),
					Year:   container["Year"].(string),
					ImdbID: container["imdbID"].(string),
					Type:   container["Type"].(string),
					Poster: container["Poster"].(string),
				})
			}
			return &pb.Response{
				Response:    "True",
				TotalResult: r["totalResults"].(string),
				Searh:       movies,
			}
		}
	}
	return &pb.Response{
		Response: "False",
		Error:    r["Error"].(string),
	}
}

type ResponseDetail struct {
	Response   string     `json:"Response"`
	Error      string     `json:"Error"`
	Title      string     `json:"Title"`
	Year       string     `json:"Year"`
	Rated      string     `json:"Rated"`
	Released   string     `json:"Released"`
	Runtime    string     `json:"Runtime"`
	Genre      string     `json:"Genre"`
	Director   string     `json:"Director"`
	Writer     string     `json:"Writer"`
	Actors     string     `json:"Actors"`
	Plot       string     `json:"Plot"`
	Language   string     `json:"Language"`
	Country    string     `json:"Country"`
	Awards     string     `json:"Awards"`
	Poster     string     `json:"Poster"`
	Ratings    []*Ratings `json:"Ratings"`
	Metascore  string     `json:"Metascore"`
	ImdbRating string     `json:"imdbRating"`
	ImdbVotes  string     `json:"imdbVotes"`
	ImdbID     string     `json:"imdbID"`
	Type       string     `json:"Type"`
	DVD        string     `json:"DVD"`
	BoxOffice  string     `json:"BoxOffice"`
	Production string     `json:"Production"`
	Website    string     `json:"Website"`
}

type Ratings struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

func (r Response) ToDetail() *pb.ResponseImdbID {
	if value, ok := r["Response"].(string); value != "" && ok {
		if strings.EqualFold(value, "true") {
			ratings := []*pb.Ratings{}
			rating := r["Ratings"].([]interface{})
			for _, v := range rating {
				container, _ := v.(map[string]interface{})
				ratings = append(ratings, &pb.Ratings{
					Source: container["Source"].(string),
					Value:  container["Value"].(string),
				})
			}
			return &pb.ResponseImdbID{
				Response:   "True",
				Actors:     r["Actors"].(string),
				Awards:     r["Awards"].(string),
				BoxOffice:  r["BoxOffice"].(string),
				Country:    r["Country"].(string),
				Director:   r["Director"].(string),
				DVD:        r["DVD"].(string),
				Genre:      r["Genre"].(string),
				ImdbID:     r["imdbID"].(string),
				ImdbRating: r["imdbRating"].(string),
				ImdbVotes:  r["imdbVotes"].(string),
				Language:   r["Language"].(string),
				Metascore:  r["Metascore"].(string),
				Plot:       r["Plot"].(string),
				Poster:     r["Poster"].(string),
				Production: r["Production"].(string),
				Rated:      r["Rated"].(string),
				Ratings:    ratings,
				Released:   r["Released"].(string),
				Runtime:    r["Runtime"].(string),
				Title:      r["Title"].(string),
				Type:       r["Type"].(string),
				Website:    r["Website"].(string),
				Writer:     r["Writer"].(string),
				Year:       r["Year"].(string),
			}
		}
	}
	return &pb.ResponseImdbID{
		Response: "False",
		Error:    r["Error"].(string),
	}
}
