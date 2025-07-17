package api

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/jose-techcode/CLI_Luasys/sysinfo"
)

func Server() {
    e := echo.New()

    e.GET("/", func(c echo.Context) error {
        informations := sysinfo.Runsys()
        return c.String(http.StatusOK, informations)
    })

    e.Logger.Fatal(e.Start(":8000"))
}