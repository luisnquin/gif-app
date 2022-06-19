//nolint:typecheck
package controllers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/luisnquin/gif-app/src/server/log"
)

func (s *ServiceMan) GetFullProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		profile, err := s.provider.GetFullProfileByUsername(c.Request().Context(), c.Param("username"))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return echo.ErrNotFound
			}

			log.Error(err)

			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, profile)
	}
}
