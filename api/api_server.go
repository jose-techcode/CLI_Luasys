package api

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/jose-techcode/CLI_Luasys/sysinfo"
)

func Server() {
    e := echo.New()

    // GET

    e.GET("/", func(c echo.Context) error {
        informations := sysinfo.Runsys()
        return c.String(http.StatusOK, informations)
    })

    e.GET("/cpu", func(c echo.Context) error {
        cpuInfo := sysinfo.GetCPUInfo()
        return c.String(http.StatusOK, cpuInfo)
    })

    e.GET("/gpu", func(c echo.Context) error {
        gpuInfo := sysinfo.GetGPUInfo()
        return c.String(http.StatusOK, gpuInfo)
    })

    e.GET("/motherboard", func(c echo.Context) error {
        motherboardInfo := sysinfo.GetMOTHERBOARDInfo()
        return c.String(http.StatusOK, motherboardInfo)
    })

    e.GET("/bios", func(c echo.Context) error {
        biosInfo := sysinfo.GetBIOSInfo()
        return c.String(http.StatusOK, biosInfo)
    })

    e.GET("/mem", func(c echo.Context) error {
        memInfo := sysinfo.GetMEMInfo()
        return c.String(http.StatusOK, memInfo)
    })

    e.GET("/disk", func(c echo.Context) error {
        diskInfo := sysinfo.GetDISKInfo()
        return c.String(http.StatusOK, diskInfo)
    })

    e.GET("/battery", func(c echo.Context) error {
        batteryInfo := sysinfo.GetBATTERYInfo()
        return c.String(http.StatusOK, batteryInfo)
    })

    e.GET("/temperature", func(c echo.Context) error {
        temperatureInfo := sysinfo.GetTEMPERATUREInfo()
        return c.String(http.StatusOK, temperatureInfo)
    })

    e.GET("/usb", func(c echo.Context) error {
        usbInfo := sysinfo.GetUSBInfo()
        return c.String(http.StatusOK, usbInfo)
    })

    e.GET("/host", func(c echo.Context) error {
        hostInfo := sysinfo.GetHOSTInfo()
        return c.String(http.StatusOK, hostInfo)
    })

    e.GET("/net", func(c echo.Context) error {
        netInfo := sysinfo.GetNETInfo()
        return c.String(http.StatusOK, netInfo)
    })

    e.GET("/processes", func(c echo.Context) error {
        processesInfo := sysinfo.GetPROCESSESInfo()
        return c.String(http.StatusOK, processesInfo)
    })

    e.Logger.Fatal(e.Start(":8000"))
}