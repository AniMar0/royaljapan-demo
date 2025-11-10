/*
Package main implements the Royal Japan Demo Backend API.

Tech Stack:
  - Language: Go 1.23+
  - Web Framework: Gin (github.com/gin-gonic/gin)
  - HTTP Server: Built-in net/http
  - Port: 8080

Why Gin Framework?
Gin was chosen for this project because it provides:
  1. High Performance: One of the fastest Go web frameworks, using httprouter under the hood
  2. Lightweight: Minimal memory footprint and fast routing
  3. Developer-Friendly: Clean API with intuitive middleware support
  4. JSON Support: Built-in JSON validation and rendering (gin.H shorthand)
  5. Production-Ready: Includes recovery middleware, logging, and error handling
  6. Active Community: Well-maintained with extensive documentation

API Endpoints:
  GET /healthz - Health check endpoint returning service status and version

Version Management:
The application supports multiple version sources with the following priority:
  1. VERSION file (plain text)
  2. pom.xml (Maven projects)
  3. package.json (Node.js projects)
  4. APP_VERSION environment variable
  5. Default fallback: "1.0.0"

Usage:
  go run main.go
  # Server starts on http://localhost:8080
  # Test: curl http://localhost:8080/healthz
*/
package main

import (
	"encoding/json"
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
)

// getAppVersion reads the application version from multiple sources in priority order.
//
// It checks the following sources sequentially and returns the first valid version found:
//  1. VERSION file - Plain text file containing version string
//  2. pom.xml - Maven configuration file (reads <version> tag)
//  3. package.json - Node.js package file (reads "version" field)
//  4. APP_VERSION - Environment variable
//  5. Default - Returns "1.0.0" if no version source is found
//
// All version strings are trimmed of leading/trailing whitespace.
//
// Returns:
//   - string: The application version (e.g., "0.1.0")
func getAppVersion() string {
	// Try to read from VERSION file
	if data, err := os.ReadFile("VERSION"); err == nil {
		version := strings.TrimSpace(string(data))
		if version != "" {
			return version
		}
	}

	// Try to read from pom.xml
	if file, err := os.Open("pom.xml"); err == nil {
		defer file.Close()
		if data, err := io.ReadAll(file); err == nil {
			var pom struct {
				Version string `xml:"version"`
			}
			if xml.Unmarshal(data, &pom) == nil && pom.Version != "" {
				return strings.TrimSpace(pom.Version)
			}
		}
	}

	// Try to read from package.json
	if file, err := os.Open("package.json"); err == nil {
		defer file.Close()
		if data, err := io.ReadAll(file); err == nil {
			var pkg struct {
				Version string `json:"version"`
			}
			if json.Unmarshal(data, &pkg) == nil && pkg.Version != "" {
				return strings.TrimSpace(pkg.Version)
			}
		}
	}

	// Fallback to env var
	if version := os.Getenv("APP_VERSION"); version != "" {
		return strings.TrimSpace(version)
	}

	// Default version
	return "1.0.0"
}

// main is the application entry point.
//
// It initializes the Gin router with default middleware (Logger and Recovery),
// loads the application version, and registers all API endpoints.
//
// The server listens on port 8080 and handles the following routes:
//   - GET /healthz: Returns JSON with status "ok" and app version
//
// Example Response from /healthz:
//   {
//     "status": "ok",
//     "version": "0.1.0"
//   }
func main() {
	// Initialize Gin router with Logger and Recovery middleware
	router := gin.Default()
	
	// Load application version from available sources
	appVersion := getAppVersion()

	// Health check endpoint - used for monitoring service availability
	// Returns HTTP 200 with JSON containing status and version information
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"version": appVersion,
		})
	})

	// Start HTTP server on port 8080
	router.Run(":8080")
}
