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

// getAppVersion reads the version from VERSION file, pom.xml, package.json, env var, or default
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

func main() {
	router := gin.Default()
	appVersion := getAppVersion()

	// Health endpoint
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"version": appVersion,
		})
	})

    router.Run(":8080") // Run the server in port 8080
}
