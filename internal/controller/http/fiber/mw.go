package fiber

import (
	"github.com/gofiber/fiber/v2"
	"okami-qstn-bnk/internal/controller/http/models"
)

// MiddleWare
// Обозначим условные роли
// super - супер админ, имеет доступ к метрикам, документации и всем методам
// admin - имеет доступ ко всем методам, может редактировать вопросы
// user - пользователь, может решать тесты
func (ctrl *Controller) MiddleWare(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")

	if authHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
			Message: "Authorization header is required",
		})
	}

	//!FIXME Псевдо запрос к сервису авторизации, если env не local + рефреш, если надо
	// POST "http://localhost:9001/auth/get_data_from_access_token"

	// !FIXME Проверка и CORS
	//var role = "admin"
	//switch {
	//case role == "admin":
	//	ctrl.app.Use(cors.New())
	//case role == "user":
	//	ctrl.app.Use(cors.New())
	//case role == "super":
	//	ctrl.app.Use(cors.New())
	//default:
	//	return ctx.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
	//		Message: "Authorization header is required",
	//	})
	//}

	// !FIXME Limiter
	// !FIXME Timout

	return nil
}
