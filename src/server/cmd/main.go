package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/luisnquin/meow-app/src/server/config"
	"github.com/luisnquin/meow-app/src/server/log"
	"github.com/luisnquin/meow-app/src/server/store"
	"github.com/luisnquin/meow-app/src/server/utils"
)

func main() {
	preLoad()

	server := echo.New()

	var result string

	row := store.DB.QueryRow("SELECT 1;")

	if err := row.Scan(&result); err != nil {
		log.Error(err)
	}

	fmt.Println(result)

	server.Logger.Fatal(server.Start(config.Server.Internal.Port))
}

func preLoad() {
	queue := utils.HandleExit()

	disconnectDb := store.Connect(store.Postgres)

	config.Load()

	queue <- func() { log.FatalWithCheck(disconnectDb()) }

	close(queue)
}
