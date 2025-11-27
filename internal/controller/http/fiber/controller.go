package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"okami-qstn-bnk/internal/service"
)

type Controller struct {
	questionSrv service.Questions
	templateSrv service.Templates
	app         *fiber.App
	logger      *zap.Logger
}

func NewController(logger *zap.Logger, qsrv service.Questions, tsrv service.Templates, fiber *fiber.App) *Controller {
	return &Controller{
		questionSrv: qsrv,
		templateSrv: tsrv,
		logger:      logger,
		app:         fiber,
	}
}

func (ctrl *Controller) ConfigureRoutes() {
	questionBank := ctrl.app.Group("/question-bank")
	{
		questions := questionBank.Group("/questions")
		{
			questions.Post("", ctrl.CreateQuestionsHandler)
			questions.Get(":id", ctrl.GetQuestionByIDHandler)
			questions.Get("", ctrl.GetQuestionsWithFiltersHandler)
			questions.Put(":id", ctrl.UpdateQuestionHandler)
			questions.Delete(":id", ctrl.DeleteQuestionHandler)
		}

		templates := questionBank.Group("/templates")
		{
			templates.Post("", ctrl.CreateTemplateHandler)
			templates.Get(":id", ctrl.GetTemplateByIDHandler)
			templates.Get("", ctrl.GetTemplatesWithFiltersHandler)
			templates.Put(":id", ctrl.UpdateTemplateHandler)
			templates.Delete(":id", ctrl.DeleteTemplateHandler)
		}
	}
}
