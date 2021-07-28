package repository

import (
	"github.com/farismfirdaus/stockbit-technical-test/microservices/logging/model"
)

type LogInterface interface {
	Insert(*model.DbLog) error
}
