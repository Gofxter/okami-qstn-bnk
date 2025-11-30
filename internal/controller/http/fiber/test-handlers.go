package fiber

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"okami-qstn-bnk/internal/controller/http/models"
)

func (ctrl *Controller) InstantiateHandler(ctx *fiber.Ctx) error {
	var req models.InstantiateRequest

	if err := ctx.BodyParser(&req); err != nil {
		ctrl.logger.Debug("can`t to parse body requests", zap.Any("body", req), zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "invalid request body",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	sessionId, _, _, err := ctrl.srv.Instantiate(context.Background(), req.TemplateId)
	if err != nil {
		ctrl.logger.Debug("can`t to generate test", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "invalid request body",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	return ctx.JSON(models.InstantiateResponse{
		SessionId: sessionId,
	})
}
