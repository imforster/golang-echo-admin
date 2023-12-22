package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo"
	"github.com/imforster/golang-echo-admin"
)

var (
	e, a   *echo.Echo
	appVer *AppVersion
)

func startServer(port string, wg *sync.WaitGroup, e *echo.Echo) {
	defer wg.Done()

	// Start the server on the specified port
	e.Start(":" + port)
}

func main() {
	appVer = &AppVersion{
		Version: Version,
	}
	fmt.Println("My Go Application")
	fmt.Println("Version:", appVer.Version)
	var wg sync.WaitGroup
	// Echo instance
	e = echo.New()
	a = echo.New()
	a.HideBanner = true
	// a.HidePort = true

	// // Routes
	e.GET("/", handler)

	a.GET("/admin/mappings", adminMappingHandler)
	a.GET("/admin/metrics", adminGetMetricsHandler)
	a.GET("/admin/info", adminInfoHandler)
	a.GET("/admin/config", adminGetConfigHandler)
	a.GET("/admin/env", adminGetEnvironmentHandler)
	a.POST("/admin/shutdown", adminPostShutdownHandler)
	a.GET("/health", healthHandler)

	// List of ports to listen on
	ports := []string{"8080", "9090"}

	handlers := []*echo.Echo{
		e, a,
	}

	// Start servers for each port
	for i, port := range ports {
		wg.Add(1)
		go startServer(port, &wg, handlers[i])
	}

	wg.Wait()
}

func handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
