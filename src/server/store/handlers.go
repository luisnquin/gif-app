package store

import (
	"bufio"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v4"
	"github.com/luisnquin/meow-app/src/server/log"
)

func HealthHandler(db Querier) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := db.Ping()
		if err != nil {
			log.Error(err)

			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.String(http.StatusOK, "database connection alive")
	}
}

func AutoMockHandler(db Querier) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Let handle tool
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

		for _, stmt := range stmts {
			_, err := db.Exec(c.Request().Context(), stmt)
			if err != nil {
				log.Error(err)
			}
		}

		return c.String(http.StatusOK, fmt.Sprintf("%d statements executed!", len(stmts)))
	}
}
