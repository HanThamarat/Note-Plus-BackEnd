package router

import (
	"os"

	"github.com/HanThamarat/Note-Plus-BackEnd/internal/handler"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App, 
	userHdl *handler.UserHandler,
	authHdl *handler.AuthHandler,
) {
	
	router := app.Group("/api/v1");

	authGroup := router.Group("/auth_service");
	authGroup.Post("/credential", authHdl.CredentialAuthenticate);

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	userGroup := router.Group("/user_management");
	userGroup.Post("/register", userHdl.RegisterUser);
}