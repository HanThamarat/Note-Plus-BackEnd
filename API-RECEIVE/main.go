package main

import (
	"log"
	"os"

	"github.com/HanThamarat/Note-Plus-BackEnd/internal/domain"
	"github.com/HanThamarat/Note-Plus-BackEnd/internal/handler"
	"github.com/HanThamarat/Note-Plus-BackEnd/internal/infrastructure/rabbitmq"
	"github.com/HanThamarat/Note-Plus-BackEnd/internal/repository"
	"github.com/HanThamarat/Note-Plus-BackEnd/internal/router"
	"github.com/HanThamarat/Note-Plus-BackEnd/internal/usecase"
	"github.com/HanThamarat/Note-Plus-BackEnd/pkg/database"
	initial "github.com/HanThamarat/Note-Plus-BackEnd/pkg/initialize"
	pkg "github.com/HanThamarat/Note-Plus-BackEnd/pkg/load-env"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)


func main() {
	pkg.LoadEnv();
	db := database.InitDB();

	err := db.AutoMigrate(
		&domain.User{},
		&domain.Organizations{},
		&domain.Role{},
		&domain.Member{},
		&domain.Project{},
	);

	if err != nil {
		panic("Could not migrate database: " + err.Error());
	}

	initial.UserInit(db);
	initial.RoleInit(db);

	rabbit, err := rabbitmq.NewRabbitClient();

    if err != nil {
        log.Fatalf("Failed to connect rabbit: %v", err)
    }


	// user manangement
	userRepo 	:= repository.NewGormUserRepository(db);
	userUsecase := usecase.NewUserUsecase(userRepo);
	userHandler := handler.NewUserHandler(userUsecase);

	// auth
	authRepo 	:= repository.NewGormAuthRepository(db);
	authUc 		:= usecase.NewAuthUsecase(authRepo);
	authHdl		:= handler.NewAuthHandler(authUc);

	// org
	orgRepo     := repository.NewGormOrgRepository(db);
	orgUc 		:= usecase.NewOrgUsecase(orgRepo);
	orgHdl 		:= handler.NewOrgHandler(orgUc);

	// member
	memberRepo	:= repository.NewGormMemberRepository(db);
	memberUc	:= usecase.NewMemberUsecase(memberRepo);
	memberHdl	:= handler.NewMemberHandler(memberUc);

	// project
	projectRepo	:= repository.NewProjectRepository(db, rabbit);
	projectUc 	:= usecase.NewProjectUsecase(projectRepo);
	projectHdl  := handler.NewProjectHandler(projectUc);

	app := fiber.New();
	app.Use(logger.New());
	app.Use(cors.New(cors.Config{
		AllowOrigins:  os.Getenv("CORS_URL"),
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, HEAD, PUT, DELETE, PATCH",
		AllowCredentials: true,
	}));

	router.SetupRoutes(
		app, 
		userHandler,
		authHdl,
		orgHdl,
		memberHdl,
		projectHdl,
	);

	log.Fatal(app.Listen(os.Getenv("Port")));
}
