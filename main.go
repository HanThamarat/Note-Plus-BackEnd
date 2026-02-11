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
	"github.com/gofiber/fiber/v2"
)

func main() {
	pkg.LoadEnv();
	db := database.InitDB();

	err := db.AutoMigrate(
		&domain.User{},
		&domain.Organizations{},
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

	//org
	orgRepo     := repository.NewGormOrgRepository(db);
	orgUc 		:= usecase.NewOrgUsecase(orgRepo);
	orgHdl 		:= handler.NewOrgHandler(orgUc);

	app := fiber.New();

	router.SetupRoutes(
		app, 
		userHandler,
		authHdl,
		orgHdl,
	);

	log.Fatal(app.Listen(os.Getenv("Port")));
}
