// Package fiber implements the HTTP controller
package fiber

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"okami-qstn-bnk/internal/service"
)

type Controller struct {
	srv    service.Service
	app    *fiber.App
	logger *zap.Logger
}

func NewController(logger *zap.Logger, srv service.Service, fiber *fiber.App) *Controller {
	return &Controller{
		srv:    srv,
		logger: logger,
		app:    fiber,
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

		questionBank.Post("tests/instantiate", ctrl.InstantiateHandler)
		//questionBank.Get("/docs/*", swagger.New())
	}
}
