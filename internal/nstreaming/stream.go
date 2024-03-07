package nstreaming

import (
	"fmt"
	"log/slog"
	"woonbeaj/L0/internal/config"

	"github.com/nats-io/stan.go"
)

func NewConnAndSub(cfg *config.Config, log *slog.Logger, srg OrderSaver) (*Stream, error) {
	const op = "nstreaming.stream.Mustload"

	eminem := Stream{}
	eminem.name = "Stream"

	err := eminem.connect(cfg, log)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	eminem.sub = &subscriber{conn: eminem.conn, name: "Subscriber", db: srg}
	eminem.sub.subscribe(log)

	return &eminem, nil
}

func (s *Stream) connect(cfg *config.Config, log *slog.Logger) error {
	const op = "nstreaming.stream.connect"

	conn, err := stan.Connect(cfg.NATSettings.ClusterId, cfg.NATSettings.ClientId, 
				stan.SetConnectionLostHandler(func(c stan.Conn, trubl error) {
					log.Info("%s: connetion lost with error: %w", s.name, trubl)
				}))
	if err != nil {
		log.Error(fmt.Sprintf("%s: failed to connect", s.name))
		return fmt.Errorf("%s: %w", op, err)
	}

	s.conn = &conn
	log.Info(fmt.Sprintf("%s: connection was successful", s.name))

	return nil
}