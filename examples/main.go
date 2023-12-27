package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/imforster/golang-echo-admin/handler"
	"github.com/labstack/echo"
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
	a := handler.New(e, &appVer.Version)

	// // Routes
	e.GET("/", helloHandler)

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

	// Wait for interrupt signal to gracefully shutdown the server
	// quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)
	// <-quit

	// fmt.Println("Server shutting down...")

	// // Create a context to wait for a graceful shutdown
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// if err := e.Shutdown(ctx); err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("Server shutdown complete")
}

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
