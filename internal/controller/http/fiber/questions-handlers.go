package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"okami-qstn-bnk/internal/controller/http/models"
	dto "okami-qstn-bnk/internal/models/dto"
	"okami-qstn-bnk/internal/pkg/types"
)

func (ctrl *Controller) CreateQuestionsHandler(ctx *fiber.Ctx) error {
	var optionsPtr *[]dto.Option
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

	if req.Options != nil {
		options := make([]dto.Option, len(*req.Options))
		for i, opt := range *req.Options {
			options[i] = dto.Option{
				Text:      opt.Text,
				IsCorrect: opt.IsCorrect,
			}
		}

		optionsPtr = &options
	}

	if err := ctrl.srv.CreateQuestion(context.Background(),
		&dto.Question{
			Role:       req.Role,
			Topic:      req.Topic,
			Type:       req.Type,
			Difficulty: req.Difficulty,
			Text:       req.Text,
		},
		optionsPtr,
	); err != nil {
		ctrl.logger.Debug("can`t to create question", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "can`t to create question",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (ctrl *Controller) GetQuestionByIDHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")
	unmarshalId, err := uuid.Parse(id)
	if err != nil {
		ctrl.logger.Debug("can`t to parse id request", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "invalid id",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	result, err := ctrl.srv.GetQuestion(context.Background(), unmarshalId)
	if err != nil {
		ctrl.logger.Debug("can`t to get question", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "can`t to get question",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	return ctx.JSON(result)
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

	result, err := ctrl.srv.GetQuestionsCollectionWithFilters(context.Background(), query.Role, query.Topic, query.Difficulty)
	if err != nil {
		ctrl.logger.Debug("can`t to get questions", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "can`t to get questions",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	return ctx.JSON(models.GetQuestionsWithFiltersResponse{Result: result})
}

func (ctrl *Controller) UpdateQuestionHandler(ctx *fiber.Ctx) error {
	req := models.UpdateQuestionRequest{Role: nil, Topic: nil, Difficulty: nil, Text: nil}

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

	result, err := ctrl.srv.UpdateQuestion(context.Background(), &dto.Question{Role: *req.Role, Topic: *req.Topic, Difficulty: *req.Difficulty, Text: *req.Text})
	if err != nil {
		ctrl.logger.Debug("can`t to update question", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "can`t to update question",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	return ctx.JSON(result)
}

func (ctrl *Controller) DeleteQuestionHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")
	unmarshalId, err := uuid.Parse(id)
	if err != nil {
		ctrl.logger.Debug("can`t to parse id request", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "invalid id",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	if err := ctrl.srv.DeleteQuestion(context.Background(), unmarshalId); err != nil {
		ctrl.logger.Debug("can`t to delete question", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "can`t to delete question",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	return ctx.SendStatus(fiber.StatusOK)
}
