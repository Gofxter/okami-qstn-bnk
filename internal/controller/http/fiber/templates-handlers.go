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

func (ctrl *Controller) CreateTemplateHandler(ctx *fiber.Ctx) error {
	var req models.CreateTemplateRequest

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

	if err := types.ValidatePurpose(req.Purpose); err != nil {
		ctrl.logger.Debug("can`t to validate purpose", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "invalid purpose",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	if err := ctrl.templateSrv.CreateTemplate(context.Background(),
		&dto.TestTemplate{Role: req.Role, Purpose: req.Purpose}); err != nil {
		ctrl.logger.Error("can`t create template", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "failed to create template",
			ErrorCode: fiber.StatusInternalServerError,
		})
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (ctrl *Controller) GetTemplateByIDHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")
	unmarshalId, err := uuid.Parse(id)
	if err != nil {
		ctrl.logger.Debug("can`t to parse id request", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "invalid id",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	result, err := ctrl.templateSrv.GetTemplate(context.Background(), unmarshalId)
	if err != nil {
		ctrl.logger.Debug("can`t to get template", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "can`t to get question",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	return ctx.JSON(result)
}

func (ctrl *Controller) GetTemplatesWithFiltersHandler(ctx *fiber.Ctx) error {
	query := models.TemplatesQueryRequest{Role: nil, Purpose: nil}

	if err := ctx.QueryParser(&query); err != nil {
		ctrl.logger.Debug("can`t to parse query request", zap.Any("query", query), zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "invalid request filters",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	if query.Role != nil {
		if err := types.ValidateRole(*query.Role); err != nil {
			ctrl.logger.Debug("can`t to validate role")
			return ctx.JSON(models.ErrorResponse{
				Message:   "invalid role",
				ErrorCode: fiber.StatusBadRequest,
			})
		}
	}

	if query.Purpose != nil {
		if err := types.ValidatePurpose(*query.Purpose); err != nil {
			ctrl.logger.Debug("can`t to validate purpose", zap.Error(err))
			return ctx.JSON(models.ErrorResponse{
				Message:   "invalid purpose",
				ErrorCode: fiber.StatusBadRequest,
			})
		}
	}

	result, err := ctrl.templateSrv.GetTemplatesCollectionWithFilters(context.Background(), query.Role, query.Purpose)
	if err != nil {
		ctrl.logger.Debug("can`t to get templates", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "can`t to get templates",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	return ctx.JSON(result)
}

func (ctrl *Controller) UpdateTemplateHandler(ctx *fiber.Ctx) error {
	req := models.UpdateTemplateRequest{Name: nil, Role: nil, Purpose: nil}

	if err := ctx.BodyParser(&req); err != nil {
		ctrl.logger.Debug("can`t to parse body requests", zap.Any("body", req), zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "invalid request body",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	if req.Role != nil {
		if err := types.ValidateRole(*req.Role); err != nil {
			ctrl.logger.Debug("can`t to validate role")
			return ctx.JSON(models.ErrorResponse{
				Message:   "invalid role",
				ErrorCode: fiber.StatusBadRequest,
			})
		}
	}

	if req.Purpose != nil {
		if err := types.ValidatePurpose(*req.Purpose); err != nil {
			ctrl.logger.Debug("can`t to validate purpose", zap.Error(err))
			return ctx.JSON(models.ErrorResponse{
				Message:   "invalid purpose",
				ErrorCode: fiber.StatusBadRequest,
			})
		}
	}

	result, err := ctrl.templateSrv.UpdateTemplate(context.Background(),
		&dto.TestTemplate{Role: *req.Role, Purpose: *req.Purpose})
	if err != nil {
		ctrl.logger.Debug("can`t to update template", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "can`t to update template",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	return ctx.JSON(result)
}

func (ctrl *Controller) DeleteTemplateHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")
	unmarshalId, err := uuid.Parse(id)
	if err != nil {
		ctrl.logger.Debug("can`t to parse id request", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "invalid id",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	if err := ctrl.templateSrv.DeleteTemplate(context.Background(), unmarshalId); err != nil {
		ctrl.logger.Debug("can`t to delete template", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "can`t to delete template",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	return ctx.SendStatus(fiber.StatusOK)
}
