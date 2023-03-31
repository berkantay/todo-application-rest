package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func RegisterRouters(app fiber.Router, service Service) {
	handle := handler{service: service}
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"status": "user service alive"})
	})
	app.Get("/user", handle.get(context.Background()))
}

type handler struct {
	service Service
}

func (h *handler) get(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Query("user_id")
		user, err := h.service.Read(context.Background(), userId)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(&fiber.Map{
				"status": false,
				"data":   "",
				"error":  err.Error(),
			})
		}
		payload, err := json.Marshal(user)
		if err != nil {
			return err
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   payload,
		})
	}

}
