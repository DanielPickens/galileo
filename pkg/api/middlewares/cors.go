/*
Copyright © 2024 DanielPickens
Author:  Daniel Pickens
Contact: Daniel.Pickens@gmail.com
*/
package middlewares

import (
	"net/http"

	"github.com/DanielPickens/galileo/pkg/clients/cors"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CORSMiddleware() echo.MiddlewareFunc {
	corsCli := cors.GetClient()
	config := corsCli.GetConfig()

	return middleware.CORSWithConfig(middleware.CORSConfig{
		// get allowed origins from env
		AllowOrigins:     []string{config.AllowOrigins},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders:     []string{"Content-Type", "Authorization", "Cookie"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Type", "Set-Cookie"},
		MaxAge:           0,
	})
}
