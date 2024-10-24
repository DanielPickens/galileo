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
	"github.com/DanielPickens/galileo/pkg/utils/constants"

	"github.com/labstack/echo/v4"
)

func Put(c echo.Context) error {

	id := c.Param("id")

	if id == "" {
		return helpers.Error(c, constants.ERROR_ID_NOT_FOUND, nil)
	}

	f := &models.CatForm{}

	if err := c.Bind(f); err != nil {
		return helpers.Error(c, constants.ERROR_BINDING_BODY, err)
	}

	if err := helpers.Validate(f); err != nil {
		return c.JSON(http.StatusBadRequest, handlers.ValidationErrors(err))
	}

	m := f.MapToModel()

	m.ID = id

	if err := m.Update(); err != nil {
		return helpers.Error(c, err, nil)
	}

	return c.JSON(http.StatusOK, handlers.Success(m.MapToForm()))

}
