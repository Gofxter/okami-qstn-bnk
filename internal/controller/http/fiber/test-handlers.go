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

	sessionId, questions, options, err := ctrl.srv.Instantiate(context.Background(), req.TemplateId)
	if err != nil {
		ctrl.logger.Debug("can`t to generate test", zap.Error(err))
		return ctx.JSON(models.ErrorResponse{
			Message:   "invalid request body",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	questionsResponse := make([]models.QuestionInInstantiateResponse, len(questions))
	for idx, question := range questions {
		var optionsForQuestion []models.OptionInQuestionInInstantiateResponse

		for indx, option := range options {
			if options[indx].QuestionId == question.Id {
				optionsForQuestion = append(optionsForQuestion, models.OptionInQuestionInInstantiateResponse{
					Text:      option.Text,
					IsCorrect: option.IsCorrect,
				})
			}
		}

		questionsResponse[idx] = models.QuestionInInstantiateResponse{
			QuestionId: question.Id,
			Type:       question.Type,
			Difficulty: question.Difficulty,
			Text:       question.Text,
			Options:    optionsForQuestion,
		}
	}

	return ctx.JSON(models.InstantiateResponse{
		SessionId: sessionId,
		Questions: questionsResponse,
	})
}
