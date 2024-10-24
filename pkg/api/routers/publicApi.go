/*
Copyright © 2024 DanielPickens
Author:  Daniel Pickens
Contact: Daniel.Pickens@gmail.com
*/
package routers

import (
	herosHandlers "github.com/DanielPickens/galileo/pkg/api/handlers/heros"
	"github.com/DanielPickens/galileo/pkg/api/handlers/errors"
	healthHandlers "github.com/DanielPickens/galileo/pkg/api/handlers/healthz"
	"github.com/DanielPickens/galileo/pkg/api/middlewares"
	"github.com/DanielPickens/galileo/pkg/clients/logger"
	"github.com/DanielPickens/galileo/pkg/config"
	"github.com/DanielPickens/galileo/pkg/utils/constants"
)

var publicApiRouter *Router

func InitPublicAPIRouter() {
	logger.Debug("Initializing public api router ...")
	publicApiRouter = &Router{}
	publicApiRouter.Name = "public API"
	publicApiRouter.Init()

	// order is important here
	// first register development middlewares
	if config.DevModeFlag {
		logger.Debug("Registering public api development middlewares ...")
		registerPublicApiDevModeMiddleware()
	}

	// next register middlwares
	logger.Debug("Registering public api middlewares ...")
	registerPublicAPIMiddlewares()

	// next register all health check routes
	logger.Debug("Registering public api health routes ...")
	registerPublicApiHealthCheckHandlers()

	// next register security related middleware
	logger.Debug("Registering public api security middlewares ...")
	registerPublicApiSecurityMiddlewares()

	// next register all routes
	logger.Debug("Registering public api public routes ...")
	registerPublicAPIRoutes()

	// finally register default fallback error handlers
	// 404 is handled here as the last route
	logger.Debug("Registering public api error handlers ...")
	registerPublicApiErrorHandlers()

	logger.Debug("Public api registration complete.")
}

func PublicAPIRouter() *Router {
	return publicApiRouter
}

func registerPublicAPIMiddlewares() {
	publicApiRouter.RegisterPreMiddleware(middlewares.SlashesMiddleware())

	publicApiRouter.RegisterMiddleware(middlewares.LoggerMiddleware())
	publicApiRouter.RegisterMiddleware(middlewares.TimeoutMiddleware())
	publicApiRouter.RegisterMiddleware(middlewares.RequestHeadersMiddleware())
	publicApiRouter.RegisterMiddleware(middlewares.ResponseHeadersMiddleware())

	if config.Feature(constants.FEATURE_GZIP).IsEnabled() {
		publicApiRouter.RegisterMiddleware(middlewares.GzipMiddleware())
	}
}

func registerPublicApiDevModeMiddleware() {
	publicApiRouter.RegisterMiddleware(middlewares.BodyDumpMiddleware())
}

func registerPublicApiSecurityMiddlewares() {
	publicApiRouter.RegisterMiddleware(middlewares.XSSCheckMiddleware())

	if config.Feature(constants.FEATURE_CORS).IsEnabled() {
		publicApiRouter.RegisterMiddleware(middlewares.CORSMiddleware())
	}

}

func registerPublicApiErrorHandlers() {
	publicApiRouter.Echo.HTTPErrorHandler = errors.AutomatedHttpErrorHandler()
	publicApiRouter.Echo.RouteNotFound("/*", errors.NotFound)
}

func registerPublicApiHealthCheckHandlers() {
	health := publicApiRouter.Echo.Group("/health")
	health.GET("/alive", healthHandlers.Index)
	health.GET("/ready", healthHandlers.Ready)
}

func registerPublicAPIRoutes() {
	heros := publicApiRouter.Echo.Group("/heros")
	heros.GET("", herosHandlers.Index)
	heros.GET("/:id", herosHandlers.Get)
	heros.POST("", herosHandlers.Post)
	heros.PUT("/:id", herosHandlers.Put)
	heros.DELETE("/:id", herosHandlers.Delete)
	// add more routes here ...
}
