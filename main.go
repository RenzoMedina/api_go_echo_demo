package main

import (
	"net/http"

	"github.com/RenzoMedina/api_go_echo_demo/storage"
	"github.com/labstack/echo"
)

func main() {
	storage.StartedMigrate()

	e := echo.New()
	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello World!!")
	})
	e.Logger.Fatal(e.Start(":8080"))

}
