package routes

import (
	_middleware "github.com/dedenfarhanhub/smart-koi-be/app/middleware"
	"github.com/dedenfarhanhub/smart-koi-be/controllers/auths"
	"github.com/dedenfarhanhub/smart-koi-be/controllers/calculate_productions"
	"github.com/dedenfarhanhub/smart-koi-be/controllers/history_productions"
	"github.com/dedenfarhanhub/smart-koi-be/controllers/profiles"
	"github.com/dedenfarhanhub/smart-koi-be/controllers/users"
	_ "github.com/dedenfarhanhub/smart-koi-be/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	swagger "github.com/swaggo/echo-swagger"
)

type ControllerList struct {
	MiddlewareLog         			_middleware.ConfigMiddleware
	JWTMiddleware         			middleware.JWTConfig
	AuthController        			auths.AuthController
	UserController        			users.UserController
	ProfileController     			profiles.ProfileController
	HistoryProductionController		history_productions.HistoryProductionController
	CalculateProductionController 	calculate_productions.CalculateProductionController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	roleAdminHeads := []string{ "ADMIN", "HEAD_OFFICER" }
	e.Use(cl.MiddlewareLog.MiddlewareLogging)

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// - swagger
	e.GET("/swagger/*", swagger.WrapHandler)

	auth := e.Group("/auth")
	auth.POST("/login", cl.AuthController.Login)

	user := e.Group("/users")
	user.POST("", cl.UserController.Store, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation(roleAdminHeads))
	user.GET("", cl.UserController.Fetch, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation(roleAdminHeads))
	user.GET("/:id", cl.UserController.FindById, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation(roleAdminHeads))
	user.PUT("/:id", cl.UserController.Update, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation(roleAdminHeads))
	user.DELETE("/:id", cl.UserController.Destroy, middleware.JWTWithConfig(cl.JWTMiddleware), _middleware.RoleValidation(roleAdminHeads))

	profile := e.Group("/profile")
	profile.GET("", cl.ProfileController.FindByToken, middleware.JWTWithConfig(cl.JWTMiddleware))
	profile.PUT("", cl.ProfileController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))

	historyProduction := e.Group("/history-productions")
	historyProduction.POST("", cl.HistoryProductionController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
	historyProduction.GET("", cl.HistoryProductionController.Fetch, middleware.JWTWithConfig(cl.JWTMiddleware))
	historyProduction.GET("/barchart", cl.HistoryProductionController.Barchart, middleware.JWTWithConfig(cl.JWTMiddleware))
	historyProduction.GET("/stat", cl.HistoryProductionController.Stat, middleware.JWTWithConfig(cl.JWTMiddleware))
	historyProduction.GET("/:id", cl.HistoryProductionController.FindById, middleware.JWTWithConfig(cl.JWTMiddleware))
	historyProduction.PUT("/:id", cl.HistoryProductionController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))
	historyProduction.DELETE("/:id", cl.HistoryProductionController.Destroy, middleware.JWTWithConfig(cl.JWTMiddleware))


	calculateProduction := e.Group("/calculate-productions")
	calculateProduction.POST("", cl.CalculateProductionController.Store, middleware.JWTWithConfig(cl.JWTMiddleware))
}