package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"

	"github.com/labstack/echo"
	"github.com/fukata/golang-stats-api-handler"
)

var (
	e, a *echo.Echo
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
	a.GET("/admin/stats", adminGetStatsHandler)
	a.GET("/admin/version", adminVersionHandler)

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

func adminVersionHandler(c echo.Context) error {
	c.JSONPretty(http.StatusOK, appVer, " ")
	return nil
}

func adminGetStatsHandler(c echo.Context) error {
	stats := stats_api.GetStats()
	c.JSONPretty(http.StatusOK, stats, " ")
	return nil
}

func adminMappingHandler(c echo.Context) error {
	// Retrieve registered routes
	routes := e.Routes()
	adminRoutes := a.Routes()
	for _, r := range adminRoutes {
		routes = append(routes, r)
	}

	// Create a new template and parse the HTML file
	t := template.Must(template.ParseFiles("admin.html"))

	// Execute the template with the routes data
	data := struct {
		Routes []*echo.Route
	}{
		Routes: routes,
	}

	// Execute the template and handle errors
	err := t.Execute(c.Response().Writer, data)
	if err != nil {
		// Handle the error, for instance, log it
		fmt.Println("Template execution error:", err)
		return err
	}

	return nil
}
