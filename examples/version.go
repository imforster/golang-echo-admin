// version.go

package main

var (
    // Version holds the current version of the application
    Version string
)

type AppVersion struct {
	Version string `json: "version"`
}
