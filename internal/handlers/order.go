package handlers

import (
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type Request struct {
	OrderId string `json:"orderId"`
}

type Response struct {
	Status  string `json:"status"`
	Error   string `json:"error,omitempty"`
	ModelJs string `json:"orderInfo,omitempty"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=OrderLoader
type OrderLoader interface {
	Get(orderUid string) ([]byte, error)
}

func Order(log *slog.Logger, srg OrderLoader) fiber.Handler {
	const op = "handlers.order.Order"
	return func(c *fiber.Ctx) error {
		log.Info("New request!")
		req := new(Request)

		if err := c.BodyParser(req); err != nil {
			log.Error(fmt.Sprintf("%s: failed to decode request", op))
			return c.Status(fiber.StatusBadRequest).JSON(Response{Status: "ERROR", Error: "failed to decode request"})
		}

		res, err := srg.Get(req.OrderId)

		if err != nil {
			log.Error(fmt.Sprintf("%s: error getting data from the database", op))
			return c.Status(fiber.StatusNotFound).JSON(Response{Status: "ERROR", Error: "error getting data from the database"})
		}
		return c.Status(fiber.StatusOK).JSON(Response{Status: "OK", ModelJs: string(res)})
	}
}
