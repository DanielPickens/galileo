/*
Copyright © 2024 DanielPickens
Author:  Daniel Pickens
Contact: Daniel.Pickens@gmail.com
*/
package healthz

import (
	"net/http"

	"github.com/DanielPickens/galileo/pkg/api/handlers"
	"github.com/DanielPickens/galileo/pkg/api/helpers"
	"github.com/DanielPickens/galileo/pkg/clients/dbc"

	"github.com/labstack/echo/v4"
)

func Ready(c echo.Context) error {
	dbClient := dbc.GetDBClient()
	if err := dbClient.Ping(); err != nil {
		return helpers.Error(c, err, nil)
	}

	payload := map[string]string{
		"message": "ready",
	}
	return c.JSON(http.StatusOK, handlers.Success(payload))
}
