package router

import (
	"github.com/HanThamarat/Note-Plus-BackEnd/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App, 
	userHdl *handler.UserHandler,
) {
	
	router := app.Group("/api/v1");

	userGroup := router.Group("/user_management");
	userGroup.Post("/register", userHdl.RegisterUser);
}