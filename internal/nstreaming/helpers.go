package nstreaming

import (
	// "errors"

	"github.com/nats-io/stan.go"
)

type OrderSaver interface {
	Save(orderUid string, orderInfo []byte) error
}

type subscriber struct {
	sub  	stan.Subscription
	db		OrderSaver
	conn 	*stan.Conn
	name 	string
}

type Stream struct {
	conn 	*stan.Conn
	sub  	*subscriber
	name 	string
}

// var ErrNotValid = errors.New("the data came in, but it was invalid") 
