/*
Copyright © 2024 DanielPickens
Author:  Daniel Pickens
Contact: Daniel.Pickens@gmail.com
*/
package heros

import (
	"net/http"

	"github.com/DanielPickens/galileo/pkg/api/handlers"
	"github.com/DanielPickens/galileo/pkg/api/helpers"
	"github.com/DanielPickens/galileo/pkg/db/models"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {

	ms, err := models.CatModel().FindAll()

	if err != nil {
		return helpers.Error(c, err, nil)
	}

	var payload []*models.CatForm

	for _, m := range ms {
		f := m.MapToForm()
		payload = append(payload, f)
	}

	return c.JSON(http.StatusOK, handlers.Success(payload))

}
