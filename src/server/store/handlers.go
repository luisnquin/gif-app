package store

import (
	"bufio"
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v4"
	"github.com/luisnquin/meow-app/src/server/log"
)

func (m *database) HealthHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := m.db.Ping()
		if err != nil {
			log.Error(err)

			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.String(http.StatusOK, "connection alive")
	}
}

func (m *database) AutoMockHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		cmd := exec.CommandContext(c.Request().Context(), "./tools/automock/main.py", "--stdout", "--length=10")
		pipe, err := cmd.StdoutPipe()
		if err != nil {
			return echo.ErrInternalServerError
		}

		defer pipe.Close()

		err = cmd.Start()
		if err != nil {
			return echo.ErrInternalServerError
		}

		stmts := make([]string, 0)

		s := bufio.NewScanner(pipe)
		for s.Scan() {
			stmts = append(stmts, s.Text())
		}

		err = cmd.Wait()
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.NoContent(http.StatusOK)
	}
}
