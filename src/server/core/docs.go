package core

import (
	"github.com/labstack/echo/v4"
	"github.com/mvrilo/go-redoc"
	echoredoc "github.com/mvrilo/go-redoc/echo"
)

func Docs() echo.MiddlewareFunc {
	doc := redoc.Redoc{
		Title:       "API Docs",
		Description: "See API Documentation",
		SpecFile:    "./docs/openapi.yaml",
		SpecPath:    "/docs/openapi.yaml",
		DocsPath:    "/docs",
	}

	return echoredoc.New(doc)
}
