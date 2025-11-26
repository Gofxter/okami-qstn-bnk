package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"okami-qstn-bnk/internal/controller/http/models"
)

func (ctrl *Controller) CreateTemplateHandler(ctx *fiber.Ctx) error {
	var req models.CreateTemplateRequest

	if err := ctx.BodyParser(&req); err != nil {
		ctrl.logger.Debug("can`t to parse body requests", zap.Any("body", req), zap.Error(err))
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (ctrl *Controller) GetTemplateByIDHandler(ctx *fiber.Ctx) error {
	_, err := ctx.ParamsInt("id")
	if err != nil {
		ctrl.logger.Debug("can`t to parse id request", zap.Error(err))
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(models.GetTemplateByIDResponse{})
}

func (ctrl *Controller) GetTemplatesWithFiltersHandler(ctx *fiber.Ctx) error {
	var query models.TemplatesQueryRequest

	if err := ctx.QueryParser(&query); err != nil {
		ctrl.logger.Debug("can`t to parse query request", zap.Any("query", query), zap.Error(err))
	}

	return ctx.JSON(models.GetTemplatesWithFiltersResponse{})
}

func (ctrl *Controller) UpdateTemplateHandler(ctx *fiber.Ctx) error {
	var req models.UpdateTemplateRequest

	if err := ctx.BodyParser(&req); err != nil {
		ctrl.logger.Debug("can`t to parse body requests", zap.Any("body", req), zap.Error(err))
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(models.UpdateTemplateResponse{})
}

func (ctrl *Controller) DeleteTemplateHandler(ctx *fiber.Ctx) error {
	_, err := ctx.ParamsInt("id")
	if err != nil {
		ctrl.logger.Debug("can`t to parse id request", zap.Error(err))
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
