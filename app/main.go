package main

import (
	_config "github.com/dedenfarhanhub/smart-koi-be/app/config"
	_dbMysqlDriver "github.com/dedenfarhanhub/smart-koi-be/drivers/mysql"
	"github.com/dedenfarhanhub/smart-koi-be/helper/logging"
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"time"

	_middleware "github.com/dedenfarhanhub/smart-koi-be/app/middleware"
	_routes "github.com/dedenfarhanhub/smart-koi-be/app/routes"
	_calculateProductionUseCase "github.com/dedenfarhanhub/smart-koi-be/business/calculate_productions"
	_historyProductionUseCase "github.com/dedenfarhanhub/smart-koi-be/business/history_productions"
	_userUseCase "github.com/dedenfarhanhub/smart-koi-be/business/users"
	_authController "github.com/dedenfarhanhub/smart-koi-be/controllers/auths"
	_calculateProductionController "github.com/dedenfarhanhub/smart-koi-be/controllers/calculate_productions"
	_historyProductionController "github.com/dedenfarhanhub/smart-koi-be/controllers/history_productions"
	_profileController "github.com/dedenfarhanhub/smart-koi-be/controllers/profiles"
	_userController "github.com/dedenfarhanhub/smart-koi-be/controllers/users"
	_calculateProductionRepo "github.com/dedenfarhanhub/smart-koi-be/drivers/databases/calculate_productions"
	_historyProductionRepo "github.com/dedenfarhanhub/smart-koi-be/drivers/databases/history_productions"
	_userRepo "github.com/dedenfarhanhub/smart-koi-be/drivers/databases/users"
)

func main()  {
	configApp := _config.GetConfig()
	mysqlConfigDB := _dbMysqlDriver.ConfigDB{
		DbUsername: configApp.Mysql.User,
		DbPassword: configApp.Mysql.Pass,
		DbHost:     configApp.Mysql.Host,
		DbPort:     configApp.Mysql.Port,
		DbDatabase: configApp.Mysql.Name,
	}
	mysqlDb := mysqlConfigDB.InitialMysqlDB()

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       configApp.JWT.Secret,
		ExpiresDuration: configApp.JWT.Expired,
	}

	timeoutContext := time.Duration(configApp.JWT.Expired) * time.Second

	e := echo.New()

	logger := logging.NewLogger()

	middlewareLog := _middleware.NewMiddleware(logger)

	userRepo := _userRepo.NewMySQLUsersRepository(mysqlDb)
	userUseCase := _userUseCase.NewUserUseCase(userRepo, &configJWT, timeoutContext, logger)
	userCtrl := _userController.NewUserController(userUseCase)

	authCtrl := _authController.NewAuthController(userUseCase)

	profileCtrl := _profileController.NewProfileController(userUseCase)

	historyProductionRepo := _historyProductionRepo.NewMySQLHistoryProductionsRepository(mysqlDb)
	historyProductionUseCase := _historyProductionUseCase.NewHistoryProductionUseCase(historyProductionRepo, timeoutContext, logger)
	historyProductionCtrl := _historyProductionController.NewHistoryProductionController(historyProductionUseCase)

	calculateProductionRepo := _calculateProductionRepo.NewMySQLCalculateProductionsRepository(mysqlDb)
	calculateProductionUseCase := _calculateProductionUseCase.NewCalculateProductionUseCase(calculateProductionRepo, historyProductionRepo, timeoutContext, logger)
	calculateProductionCtrl := _calculateProductionController.NewCalculateProductionController(calculateProductionUseCase)

	routesInit := _routes.ControllerList{
		MiddlewareLog:         			middlewareLog,
		JWTMiddleware:         			configJWT.Init(),
		AuthController:        			*authCtrl,
		UserController:        			*userCtrl,
		ProfileController: 	   			*profileCtrl,
		HistoryProductionController:	*historyProductionCtrl,
		CalculateProductionController: 	*calculateProductionCtrl,
	}
	routesInit.RouteRegister(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Print("listening on PORT : ", port)
	log.Fatal(e.Start(":" + port))
}
