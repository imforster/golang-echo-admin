package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	stats_api "github.com/fukata/golang-stats-api-handler"
	"github.com/labstack/echo"
)

var (
	e, a   *echo.Echo
	appVer *string
	path   string
)

func New(mainEcho *echo.Echo, version *string, templatePath string) *echo.Echo {
	e = mainEcho
	adminEcho := echo.New()
	adminEcho.HideBanner = true
	a = adminEcho
	if version != nil {
		appVer = version
	}
	path = templatePath

	a.GET("/admin/mappings", adminMappingHandler)
	a.GET("/admin/metrics", AdminGetMetricsHandler)
	a.GET("/admin/info", AdminInfoHandler)
	a.GET("/admin/config", AdminGetConfigHandler)
	a.GET("/admin/env", adminGetEnvironmentHandler)
	a.POST("/admin/shutdown", AdminPostShutdownHandler)
	a.GET("/health", HealthHandler)
	return adminEcho
}

func adminGetEnvironmentHandler(c echo.Context) error {
	// Get all environment variables
	envVariables := make(map[string]string)
	for _, env := range os.Environ() {
		keyValue := strings.SplitN(env, "=", 2)
		envVariables[keyValue[0]] = keyValue[1]
	}

	// Convert map to a pretty-printed JSON string
	prettyJSON, err := json.MarshalIndent(envVariables, "", "  ")
	if err != nil {
		return err
	}

	// Return the pretty-printed JSON as a response
	return c.JSONBlob(http.StatusOK, prettyJSON)
}

func AdminPostShutdownHandler(c echo.Context) error {
	go func() {
		c.Echo().Logger.Info("shutting down the server")

		// Graceful shutdown with a timeout of 10 seconds
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := e.Shutdown(context.Background())
		if err != nil {
			if err != http.ErrServerClosed {
				c.Echo().Logger.Fatal("shutting down the server")
			}
		}

		err = a.Shutdown(context.Background())
		if err != nil {
			if err != http.ErrServerClosed {
				c.Echo().Logger.Fatal("shutting down the server")
			}
		}
	}()

	return c.String(http.StatusOK, "Shutting down..")
}

func HealthHandler(c echo.Context) error {
	// Perform a simple check, like sending a request to root endpoint
	resp, err := http.Get("http://localhost:8080")
	if err != nil || resp.StatusCode != http.StatusOK {
		return c.String(http.StatusInternalServerError, "Echo server is not responsive")
	}

	// You can perform other checks here if needed

	// If all checks pass, indicate the server is healthy
	status := map[string]string{
		"status": "UP",
	}
	prettyJSON, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		return err
	}
	return c.JSONBlob(http.StatusOK, prettyJSON)
}

func AdminInfoHandler(c echo.Context) error {
	c.JSONPretty(http.StatusOK, appVer, " ")
	return nil
}

func AdminGetMetricsHandler(c echo.Context) error {
	stats := stats_api.GetStats()
	c.JSONPretty(http.StatusOK, stats, " ")
	return nil
}

func AdminGetConfigHandler(c echo.Context) error {
	config := "TBD config stuff"
	c.JSONPretty(http.StatusOK, config, " ")
	return nil
}

func adminMappingHandler(c echo.Context) error {
	// Retrieve registered routes
	routes := e.Routes()
	adminRoutes := a.Routes()

	// Check Accept header for response format preference
	if c.Request().Header.Get("Accept") == "application/json" {
		for _, r := range adminRoutes {
			routes = append(routes, r)
		}
		// Return routes as JSON response
		return c.JSONPretty(http.StatusOK, routes, " ")
	}

	// Create a new template and parse the HTML file
	adminPath := filepath.Join(path, "admin.html")
	// t := template.Must(template.ParseFiles("../handler/admin.html"))
	t := template.Must(template.ParseFiles(adminPath))

	// Execute the template with the routes data
	data := struct {
		Routes      []*echo.Route
		AdminRoutes []*echo.Route
		Port        string
		AdminPort   string
	}{
		Routes:      routes,
		AdminRoutes: adminRoutes,
		Port:        e.Server.Addr,
		AdminPort:   a.Server.Addr,
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
