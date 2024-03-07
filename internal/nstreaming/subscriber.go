package nstreaming

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"woonbeaj/L0/internal/jsonStruct"

	"github.com/nats-io/stan.go"
)

func (s *subscriber) subscribe(log *slog.Logger) error {
	const op = "nstreaming.subscriber.subscribe"

	var err error
	s.sub, err = (*s.conn).Subscribe(
									"L0_enjoyer",
									s.newMessage(log),
									stan.AckWait(stan.DefaultAckWait),
									stan.SetManualAckMode(),
									)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info(fmt.Sprintf("%s: contract signed", s.name))
	return nil
}

func(s *subscriber) newMessage(log *slog.Logger) stan.MsgHandler{
	return func(msg *stan.Msg) {
		log.Info("new message received")
		defer func() {
			err := msg.Ack()
			if err != nil {
				log.Error(fmt.Sprintf("%s: failed to acknowledge message", s.name))
			}
		}()
		var order jsonStruct.OrderInfo

		err := json.Unmarshal(msg.Data, &order)

		if err != nil {
			log.Error(fmt.Sprintf("%s: the data came in, but it was invalid: %v", s.name, err))
			return
		}

		err = s.db.Save(order.OrderUID, msg.Data)
		
		if err != nil {
			log.Error(fmt.Sprintf("%s: unable to add order: %v\n", s.name, err))
			return 
		}
	}
}

func (s *subscriber) Unsubscribe() {
	if s.sub != nil {
		s.sub.Unsubscribe()
	}
}