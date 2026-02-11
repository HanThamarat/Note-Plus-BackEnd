package router

import (
	"os"

	"github.com/HanThamarat/Note-Plus-BackEnd/internal/handler"
	"github.com/HanThamarat/Note-Plus-BackEnd/pkg/responses"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func SetupRoutes(
	app *fiber.App, 
	userHdl *handler.UserHandler,
	authHdl *handler.AuthHandler,
	orgHdl 	*handler.OrgHandler,
) {
	app.Get("/", func (c *fiber.Ctx) error {
		return responses.SetResponse(c, fiber.StatusOK, "Server is runing", nil); 
	});

	router := app.Group("/api/v1");

	authGroup := router.Group("/auth_service");
	authGroup.Post("/credential", authHdl.CredentialAuthenticate);

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	userGroup := router.Group("/user_management");
	userGroup.Post("/register", userHdl.RegisterUser);

	// org management
	orgGruop := router.Group("/org_service");
	orgGruop.Post("/organization", orgHdl.CreateNewOrg);
	orgGruop.Get("/organization", orgHdl.FindAllOrg);
	orgGruop.Get("/organization/:id", orgHdl.FindOrgById);
	orgGruop.Put("/organization/:id", orgHdl.UpdateOrg);
	orgGruop.Delete("/organization/:id", orgHdl.DeleteOrg);
	
}