package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"okami-qstn-bnk/internal/controller/http/models"
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

	return ctx.SendStatus(fiber.StatusOK)
}

func (ctrl *Controller) GetTemplateByIDHandler(ctx *fiber.Ctx) error {
	_, err := ctx.ParamsInt("id")
	if err != nil {
		ctrl.logger.Debug("can`t to parse id request", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "invalid id",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	return ctx.JSON(models.GetTemplateByIDResponse{})
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

	return ctx.JSON(models.GetTemplatesWithFiltersResponse{})
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

	return ctx.JSON(models.UpdateTemplateResponse{})
}

func (ctrl *Controller) DeleteTemplateHandler(ctx *fiber.Ctx) error {
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
