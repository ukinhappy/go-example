package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type res struct {
	Code int
	Msg  string
	Data map[string]interface{}
}

func main() {
	e := echo.New()
	e.GET("/userinfo", func(ctx echo.Context) error {
		uid := ctx.QueryParam("uid")
		fmt.Println(uid)
		if uid == "" {
			return ctx.String(http.StatusOK, "缺少uid参数")
		}
		return ctx.JSON(http.StatusOK, res{Code: 200, Msg: "success", Data: map[string]interface{}{"name": "zdx"}})
	})

	e.POST("/say", func(ctx echo.Context) error {
		uid := ctx.QueryParam("uid")
		if uid == "" {
			return ctx.String(http.StatusOK, "缺少uid参数")
		}
		return ctx.JSON(http.StatusOK, res{Code: 200, Msg: "success", Data: map[string]interface{}{"text": fmt.Sprintf("你好%s", uid)}})
	})
	e.Start("0.0.0.0:8080")
}
