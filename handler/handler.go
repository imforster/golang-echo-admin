package handler

import (
	"fmt"
	"html/template"
	"net/http"

	stats_api "github.com/fukata/golang-stats-api-handler"
	"github.com/labstack/echo"
)

var (
	e, a   *echo.Echo
	appVer *string
)

func Handler(version *string) {
	appVer = version
}

func adminGetEnvironmentHandler(c echo.Context) error {
	env := "Environment stuff..."
	return c.String(http.StatusOK, env)
}

func adminPostShutdownHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Shutting down..")
}

func healthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "{status = \"UP\"}")
}

func adminInfoHandler(c echo.Context) error {
	c.JSONPretty(http.StatusOK, appVer, " ")
	return nil
}

func adminGetMetricsHandler(c echo.Context) error {
	stats := stats_api.GetStats()
	c.JSONPretty(http.StatusOK, stats, " ")
	return nil
}

func adminGetConfigHandler(c echo.Context) error {
	config := "TBD config stuff"
	c.JSONPretty(http.StatusOK, config, " ")
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
