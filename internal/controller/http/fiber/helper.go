package fiber

import "github.com/gofiber/fiber/v2"

func (ctrl *Controller) BaseHandler(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}
