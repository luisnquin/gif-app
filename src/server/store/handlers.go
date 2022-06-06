package store

import (
	"bufio"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/luisnquin/gif-app/src/server/log"
)

func HealthHandler(db Querier, cache *redis.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := cache.Ping(c.Request().Context()).Err()
		if err != nil {
			log.Error(err)

			return c.String(http.StatusInternalServerError, err.Error())
		}

		err = db.Ping()
		if err != nil {
			log.Error(err)

			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.String(http.StatusOK, "database and cache connection alive\n")
	}
}

func AutoMockHandler(db Querier) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Let handle tool
		cmd := exec.CommandContext(c.Request().Context(), "python3", "./tools/automock/main.py", "--stdout", "--length=10")
		pipe, err := cmd.StdoutPipe()
		if err != nil {
			log.Error(err)

			return echo.ErrInternalServerError
		}

		defer pipe.Close()

		err = cmd.Start()
		if err != nil {
			log.Error(err)

			return echo.ErrInternalServerError
		}

		stmts := make([]string, 0)

		s := bufio.NewScanner(pipe)
		for s.Scan() {
			stmts = append(stmts, s.Text())
		}

		err = cmd.Wait()
		if err != nil {
			log.Error(err)

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
