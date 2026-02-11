package main

import (
	"log"
	"os"

	"github.com/HanThamarat/Note-Plus-BackEnd/internal/domain"
	"github.com/HanThamarat/Note-Plus-BackEnd/internal/handler"
	"github.com/HanThamarat/Note-Plus-BackEnd/internal/repository"
	"github.com/HanThamarat/Note-Plus-BackEnd/internal/router"
	"github.com/HanThamarat/Note-Plus-BackEnd/internal/usecase"
	"github.com/HanThamarat/Note-Plus-BackEnd/pkg/database"
	pkg "github.com/HanThamarat/Note-Plus-BackEnd/pkg/load-env"
	"github.com/HanThamarat/Note-Plus-BackEnd/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

func main() {
	pkg.LoadEnv();
	db := database.InitDB();

	err := db.AutoMigrate(
		&domain.User{},
	);

	if err != nil {
		panic("Could not migrate database: " + err.Error());
	}

	// user manangement
	userRepo 	:= repository.NewGormUserRepository(db);
	userUsecase := usecase.NewUserUsecase(userRepo);
	userHandler := handler.NewUserHandler(userUsecase);

	// auth
	authRepo 	:= repository.NewGormAuthRepository(db);
	authUc 		:= usecase.NewAuthUsecase(authRepo);
	authHdl		:= handler.NewAuthHandler(authUc);

	app := fiber.New();

	router.SetupRoutes(
		app, 
		userHandler,
		authHdl,
	);

	app.Get("", func (c *fiber.Ctx) error {
		return responses.SetResponse(c, fiber.StatusOK, "Server is runing", nil); 
	});

	log.Fatal(app.Listen(os.Getenv("Port")));
}
