package routes

import (
	"github.com/mesxx/Fiber_Task_Management_API/handlers"
	"github.com/mesxx/Fiber_Task_Management_API/middlewares"
	"github.com/mesxx/Fiber_Task_Management_API/repositories"
	"github.com/mesxx/Fiber_Task_Management_API/usecases"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TaskRoute(router fiber.Router, db *gorm.DB) {
	repository := repositories.NewTaskRepositoy(db)
	usecase := usecases.NewTaskUsecase(repository)
	handler := handlers.NewTaskHandler(usecase)

	router.Get("/", handler.GetAll)
	router.Delete("/delete/all", handler.DeleteAll)

	// Authorization
	router.Use(middlewares.RestrictedUser)

	router.Static("/image", "./publics/images")

	router.Post("/", handler.Create)
	router.Get("/user", handler.GetAllByUser)
	router.Get("/:id", handler.GetByID)
	router.Get("/image/:id", handler.GetImageByID)
	router.Patch("/:id", handler.UpdateByID)
	router.Delete("/image/:id", handler.DeleteImageByID)
	router.Delete("/:id", handler.DeleteByID)
	//
}
