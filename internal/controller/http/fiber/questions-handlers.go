package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"okami-qstn-bnk/internal/controller/http/models"
	"okami-qstn-bnk/internal/pkg/types"
)

func (ctrl *Controller) CreateQuestionsHandler(ctx *fiber.Ctx) error {
	var req models.CreateQuestionRequest

	if err := ctx.BodyParser(&req); err != nil {
		ctrl.logger.Debug("can`t to parse body requests", zap.Any("body", req), zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "invalid request body",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	if err := types.ValidateRole(req.Role); err != nil {
		ctrl.logger.Debug("can`t to validate role", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "invalid role",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	if err := types.ValidateType(req.Type); err != nil {
		ctrl.logger.Debug("can`t to validate type", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "invalid type",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (ctrl *Controller) GetQuestionByIDHandler(ctx *fiber.Ctx) error {
	_, err := ctx.ParamsInt("id")
	if err != nil {
		ctrl.logger.Debug("can`t to parse id request", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "invalid id",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	return ctx.JSON(models.GetQuestionByIDResponse{})
}

func (ctrl *Controller) GetQuestionsWithFiltersHandler(ctx *fiber.Ctx) error {
	query := models.QuestionsQueryRequest{Role: nil, Topic: nil, Difficulty: nil}

	if err := ctx.QueryParser(&query); err != nil {
		ctrl.logger.Debug("can`t to parse query request", zap.Any("query", query), zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "invalid request filters",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	if query.Role != nil {
		if err := types.ValidateRole(*query.Role); err != nil {
			ctrl.logger.Debug("can`t to validate role", zap.Error(err))
			return ctx.JSON(models.ErrorResponse{
				Message:   "invalid role",
				ErrorCode: fiber.StatusBadRequest,
			})
		}
	}

	return ctx.JSON(models.GetQuestionsWithFiltersResponse{})
}

func (ctrl *Controller) UpdateQuestionHandler(ctx *fiber.Ctx) error {
	req := models.UpdateQuestionRequest{Role: nil, Topic: nil, Type: nil, Options: nil, Difficulty: nil, Text: nil}

	if err := ctx.BodyParser(&req); err != nil {
		ctrl.logger.Debug("can`t to parse body requests", zap.Any("body", req), zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "invalid request body",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	if req.Role != nil {
		if err := types.ValidateRole(*req.Role); err != nil {
			ctrl.logger.Debug("can`t to validate role", zap.Error(err))
			return ctx.JSON(models.ErrorResponse{
				Message:   "invalid role",
				ErrorCode: fiber.StatusBadRequest,
			})
		}
	}

	if req.Type != nil {
		if err := types.ValidateType(*req.Type); err != nil {
			ctrl.logger.Debug("can`t to validate type", zap.Error(err))
			return ctx.JSON(models.ErrorResponse{
				Message:   "invalid type",
				ErrorCode: fiber.StatusBadRequest,
			})
		}
	}

	return ctx.JSON(models.UpdateQuestionResponse{})
}

func (ctrl *Controller) DeleteQuestionHandler(ctx *fiber.Ctx) error {
	_, err := ctx.ParamsInt("id")
	if err != nil {
		ctrl.logger.Debug("can`t to parse id request", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "invalid id",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	return ctx.SendStatus(fiber.StatusOK)
}
