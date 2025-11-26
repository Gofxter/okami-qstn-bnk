package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"okami-qstn-bnk/internal/controller/http/models"
)

func (ctrl *Controller) CreateQuestionsHandler(ctx *fiber.Ctx) error {
	var req models.CreateQuestionRequest

	if err := ctx.BodyParser(&req); err != nil {
		ctrl.logger.Debug("can`t to parse body requests", zap.Any("body", req), zap.Error(err))
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (ctrl *Controller) GetQuestionByIDHandler(ctx *fiber.Ctx) error {
	_, err := ctx.ParamsInt("id")
	if err != nil {
		ctrl.logger.Debug("can`t to parse id request", zap.Error(err))
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(models.GetQuestionByIDResponse{})
}

func (ctrl *Controller) GetQuestionsWithFiltersHandler(ctx *fiber.Ctx) error {
	var query models.QuestionsQueryRequest

	if err := ctx.QueryParser(&query); err != nil {
		ctrl.logger.Debug("can`t to parse query request", zap.Any("query", query), zap.Error(err))
	}

	return ctx.JSON(models.GetQuestionsWithFiltersResponse{})
}

func (ctrl *Controller) UpdateQuestionHandler(ctx *fiber.Ctx) error {
	var req models.UpdateQuestionRequest

	if err := ctx.BodyParser(&req); err != nil {
		ctrl.logger.Debug("can`t to parse body requests", zap.Any("body", req), zap.Error(err))
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.JSON(models.UpdateQuestionResponse{})
}

func (ctrl *Controller) DeleteQuestionHandler(ctx *fiber.Ctx) error {
	_, err := ctx.ParamsInt("id")
	if err != nil {
		ctrl.logger.Debug("can`t to parse id request", zap.Error(err))
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
