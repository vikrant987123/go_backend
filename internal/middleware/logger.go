package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Logger(log *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		duration := time.Since(start)

		log.Info("http_request",
			zap.String("method",c.Method()),
			zap.String("path",c.Path()),
			zap.Int("status",c.Response().StatusCode()),
			zap.Duration("duration",duration),
			zap.Any("request_id", c.Locals("request_id")),
		)

		return err
	}
}