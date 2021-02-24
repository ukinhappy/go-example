package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(ctx echo.Context) error {
		ctx.String(http.StatusOK, "111")
		return nil
	})
	e.Start("0.0.0.0:8081")
}
