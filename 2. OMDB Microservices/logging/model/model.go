package model

type DbLog struct {
	ID       int
	Method   string
	Request  string
	Response string
}
